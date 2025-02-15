use anyhow::Result;
use arcstr::ArcStr;
use hashbrown::{HashMap, HashSet};
use hcl::{
    eval::{Context, Evaluate},
    Expression,
};
use serde_json::Value as JsonValue;
use std::str::FromStr;
use std::sync::Arc;
use strum::EnumString as StrumEnumString;

use super::typ::Type;

#[inline]
fn is_null_expression(expr: &Expression) -> bool {
    match expr {
        Expression::Null => true,
        _ => false,
    }
}

#[derive(Clone, Debug)]
pub struct OverrideDecl {
    pub ident: ArcStr,
    pub body: OverrideBodyStmt,
}

#[derive(Clone, Debug)]
pub struct OverrideBodyStmt {
    pub value: Expression,
    pub context: Expression,
}

#[derive(Clone, Debug)]
pub struct ParameterDecl {
    pub ident: ArcStr,
    pub body: ParameterBodyStmt,
}

#[derive(Clone, Debug)]
pub struct ParameterBodyStmt {
    pub value_type: Expression,
    pub maximum: Expression,
    pub minimum: Expression,
    pub default: Expression,
    pub order: Expression,
    pub scope: Expression,
    pub description: Expression,
    pub excluded: Expression,
    pub protected: Expression,
}

#[derive(Clone, Debug)]
pub struct ConfigurationBodyStmt {
    pub display_name: Expression,
    pub description: Expression,
    pub order: Expression,
    pub parameters: Vec<ParameterDecl>,
    pub overrides: Vec<OverrideDecl>,
}

#[derive(Clone, Debug)]
pub enum ConfigurationDecl {
    Genesis {
        ident: ArcStr,
        body: ConfigurationBodyStmt,
    },
    Successor {
        ident: ArcStr,
        parent_ident: ArcStr,
        body: ConfigurationBodyStmt,
    },
    Anonymous {
        body: ConfigurationBodyStmt,
    },
}

impl ConfigurationDecl {
    pub fn ident(&self) -> Option<ArcStr> {
        match self {
            ConfigurationDecl::Genesis { ident, .. } => Some(ArcStr::clone(ident)),
            ConfigurationDecl::Successor { ident, .. } => Some(ArcStr::clone(ident)),
            ConfigurationDecl::Anonymous { .. } => None,
        }
    }

    pub fn parent_ident(&self) -> Option<ArcStr> {
        match self {
            ConfigurationDecl::Genesis { .. } => None,
            ConfigurationDecl::Successor { parent_ident, .. } => Some(ArcStr::clone(parent_ident)),
            ConfigurationDecl::Anonymous { .. } => None,
        }
    }

    fn body(&self) -> &ConfigurationBodyStmt {
        match self {
            ConfigurationDecl::Genesis { body, .. } => body,
            ConfigurationDecl::Successor { body, .. } => body,
            ConfigurationDecl::Anonymous { body } => body,
        }
    }
}

impl ConfigurationDecl {
    pub fn evaluate(self, ctx: &Context) -> Result<ConfigurationNode> {
        let body = self.body();
        let mut parameters = HashMap::new();

        for parameter_decl in &body.parameters {
            let typ = match Type::try_from(&parameter_decl.body.value_type) {
                Ok(ty) => ty,
                Err(_) => {
                    // TODO: Add logging for encountering an unknown type
                    continue;
                }
            };

            parameters.insert(
                parameter_decl.ident.clone(),
                Arc::new(Parameter {
                    ident: ArcStr::clone(&parameter_decl.ident),
                    typ,
                    maximum: try_evaluate_to_u64(ctx, &parameter_decl.body.maximum)?,
                    minimum: try_evaluate_to_u64(ctx, &parameter_decl.body.minimum)?,
                    default: serde_json::from_str(
                        parameter_decl
                            .body
                            .default
                            .evaluate(ctx)?
                            .to_string()
                            .as_str(),
                    )?,
                    scope: try_evaluate_to_string(ctx, &parameter_decl.body.scope)?
                        .and_then(|value| ParameterScope::from_str(&value).ok())
                        .unwrap_or_default(),
                    order: try_evaluate_to_u64(ctx, &parameter_decl.body.order)?,
                    description: try_evaluate_to_string(ctx, &parameter_decl.body.description)?,
                    excluded: try_evaluate_to_bool(ctx, &parameter_decl.body.excluded)?
                        .unwrap_or(false),
                    protected: try_evaluate_to_bool(ctx, &parameter_decl.body.protected)?
                        .unwrap_or(false),
                }),
            );
        }

        let mut overrides = HashMap::new();
        for override_decl in &body.overrides {
            let value = if is_null_expression(&override_decl.body.value) {
                // TODO: Add logging
                continue;
            } else {
                serde_json::from_str(override_decl.body.value.evaluate(ctx)?.to_string().as_str())?
            };

            let _context = if !is_null_expression(&override_decl.body.context) {
                unimplemented!()
            };

            overrides.insert(
                override_decl.ident.clone(),
                Arc::new(Override {
                    ident: ArcStr::clone(&override_decl.ident),
                    value,
                    context: None,
                }),
            );
        }

        Ok(ConfigurationNode {
            ident: self.ident().unwrap_or_default(),
            parent_ident: self.parent_ident(),
            display_name: try_evaluate_to_string(ctx, &body.display_name)?,
            description: try_evaluate_to_string(ctx, &body.description)?,
            order: try_evaluate_to_u64(ctx, &body.order)?,
            parameters,
            overrides,
        })
    }
}

fn try_evaluate_to_string(ctx: &Context, expr: &Expression) -> Result<Option<String>> {
    Ok(expr.evaluate(ctx)?.as_str().map(ToString::to_string))
}

fn try_evaluate_to_u64(ctx: &Context, expr: &Expression) -> Result<Option<u64>> {
    Ok(expr.evaluate(ctx)?.as_u64())
}

fn try_evaluate_to_bool(ctx: &Context, expr: &Expression) -> Result<Option<bool>> {
    Ok(expr.evaluate(ctx)?.as_bool())
}

#[derive(Debug, Default, StrumEnumString)]
pub enum ParameterScope {
    APPLICATION,
    #[default]
    WINDOW,
    RESOURCE,
    #[allow(non_camel_case_types)]
    LANGUAGE_SPECIFIC,
}

#[derive(Clone, Debug)]
pub struct ConfigurationNode {
    pub ident: ArcStr,
    pub parent_ident: Option<ArcStr>,
    pub display_name: Option<String>,
    pub description: Option<String>,
    pub order: Option<u64>,
    pub parameters: HashMap<ArcStr, Arc<Parameter>>,
    pub overrides: HashMap<ArcStr, Arc<Override>>,
}

#[derive(Debug)]
pub struct Parameter {
    pub ident: ArcStr,
    pub typ: Type,
    pub maximum: Option<u64>,
    pub minimum: Option<u64>,
    pub default: JsonValue,
    /// The order in which the parameter appears within its group in the settings UI.
    pub order: Option<u64>,
    pub scope: ParameterScope,
    pub description: Option<String>,
    /// Excluded parameters are hidden from the UI but can still be registered.
    pub excluded: bool,
    /// Indicates if this setting is protected from addon overrides.
    pub protected: bool,
}

#[derive(Debug)]
pub struct Override {
    pub ident: ArcStr,
    pub value: JsonValue,
    pub context: Option<HashSet<String>>,
}
