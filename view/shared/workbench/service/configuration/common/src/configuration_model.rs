use std::sync::Arc;

use platform_configuration::configuration_model::{
    AttributeName, Configuration as PlatformConfiguration, ConfigurationModel,
};
use platform_configuration::configuration_parser::ConfigurationParser as PlatformConfigurationParser;
use serde_json::Value;

pub struct WorkspaceConfigurationParser<'a> {
    platform_parser: PlatformConfigurationParser<'a>,
}

pub struct WorkspaceConfiguration {
    platform_configuration: PlatformConfiguration,
}

impl WorkspaceConfiguration {
    pub fn new(
        default_configuration: Arc<ConfigurationModel>,
        policy_configuration: Arc<ConfigurationModel>,
        user_configuration: Arc<ConfigurationModel>,
        workspace_configuration: Arc<ConfigurationModel>,
        inmem_configuration: Arc<ConfigurationModel>,
    ) -> Self {
        Self {
            platform_configuration: PlatformConfiguration::new(
                default_configuration,
                policy_configuration,
                user_configuration,
                workspace_configuration,
                inmem_configuration,
            ),
        }
    }

    pub fn get_value(&self, attribute_name: &AttributeName) -> Option<Value> {
        self.platform_configuration.get_value(attribute_name)
    }
}
