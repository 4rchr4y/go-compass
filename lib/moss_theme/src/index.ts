// FIXME: isDefault
const styleKeywords = ["background", "hoverBackground", "activeBackground", "text"];

export class Theme {
  constructor(
    public name: string,
    public type: string,
    public isDefault: boolean,
    public colors: Colors
  ) {}
}

export class Colors {
  constructor(
    public primary?: string,
    public sideBarBackground?: string,
    public toolBarBackground?: string,
    public pageBackground?: string,
    public statusBarBackground?: string,
    public windowsCloseButtonBackground?: string,
    public windowControlsLinuxBackground?: string,
    public windowControlsLinuxText?: string,
    public windowControlsLinuxHoverBackground?: string,
    public windowControlsLinuxActiveBackground?: string
  ) {}
}

// Mapping between Theme JSON and TypeScript types
const typeMap: any = {
  Theme: createPropsWithAdditional(
    Object.keys(new Theme("", "", false, new Colors())).map((key) => ({
      json: key,
      js: key,
      typ: key === "colors" ? createReference("Colors") : key === "isDefault" ? true : "",
    })),
    false
  ),
  Colors: createPropsWithAdditional(
    Object.keys(new Colors()).map((key) => {
      return {
        json: createJsonKey(key),
        js: key,
        typ: createUnionMembers(undefined, ""),
      };
    }),
    false
  ),
};

// Theme CSS variables
export type ThemeCssVariables = {
  [K in keyof Colors as `--color-${KebabCase<K>}`]: string;
};

// Map Theme to CSS variables
export function mapThemeToCssVariables(theme: Theme): ThemeCssVariables {
  const colorKeys = Object.keys(theme.colors) as (keyof Colors)[];

  return colorKeys.reduce((acc, key) => {
    const cssKey = `--color-${toKebabCase(key)}` as keyof ThemeCssVariables;
    acc[cssKey] = theme.colors[key] || "";
    return acc;
  }, {} as ThemeCssVariables);
}

// Theme custom Tailwind color variables
export const customTailwindColorVariables: Record<keyof Colors, string> = Object.keys(new Colors()).reduce(
  (acc, key) => {
    const cssKey = `--color-${toKebabCase(key)}` as keyof ThemeCssVariables;
    acc[key as keyof Colors] = rgbaWithOpacity(cssKey);
    return acc;
  },
  {} as Record<keyof Colors, string>
);

function createJsonKey(key: string) {
  const matchedKeywords = styleKeywords.filter((v) => key.toLowerCase().indexOf(v.toLowerCase()) !== -1);
  if (matchedKeywords.length > 0) {
    const longestKeyword = matchedKeywords.reduce((a, b) => (a.length > b.length ? a : b));
    return key.replace(new RegExp(longestKeyword, "i"), "." + longestKeyword);
  }
  return key;
}

type KebabCase<T extends string> = T extends `${infer F}${infer R}`
  ? R extends Uncapitalize<R>
    ? `${Lowercase<F>}${KebabCase<R>}`
    : `${Lowercase<F>}-${KebabCase<R>}`
  : T;

function toKebabCase(str: string): string {
  return str.replace(/[A-Z]/g, (letter) => `-${letter.toLowerCase()}`);
}

// https://tailwindcss.com/docs/customizing-colors#using-css-variables
function rgbaWithOpacity(variableName: keyof ThemeCssVariables): string {
  return `rgba(var(${variableName}))`;
}

export class Convert {
  public static toTheme(json: string): Theme {
    return cast(JSON.parse(json), createReference("Theme"));
  }

  public static themeToJson(value: Theme): string {
    return JSON.stringify(uncast(value, createReference("Theme")), null, 2);
  }
}

function invalidValue(typ: any, val: any, key: any, parent: any = ""): never {
  const prettyTyp = prettyTypeName(typ);
  const parentText = parent ? ` on ${parent}` : "";
  const keyText = key ? ` for key "${key}"` : "";
  throw Error(`Invalid value${keyText}${parentText}. Expected ${prettyTyp} but got ${JSON.stringify(val)}`);
}

function prettyTypeName(typ: any): string {
  if (Array.isArray(typ)) {
    if (typ.length === 2 && typ[0] === undefined) {
      return `an optional ${prettyTypeName(typ[1])}`;
    } else {
      return `one of [${typ
        .map((a) => {
          return prettyTypeName(a);
        })
        .join(", ")}]`;
    }
  } else if (typeof typ === "object" && typ.literal !== undefined) {
    return typ.literal;
  } else {
    return typeof typ;
  }
}

function jsonToJSProps(typ: any): any {
  if (typ.jsonToJS === undefined) {
    const map: any = {};
    typ.props.forEach((p: any) => (map[p.json] = { key: p.js, typ: p.typ }));
    typ.jsonToJS = map;
  }
  return typ.jsonToJS;
}

function jsToJSONProps(typ: any): any {
  if (typ.jsToJSON === undefined) {
    const map: any = {};
    typ.props.forEach((p: any) => (map[p.js] = { key: p.json, typ: p.typ }));
    typ.jsToJSON = map;
  }
  return typ.jsToJSON;
}

function transform(val: any, typ: any, getProps: any, key: any = "", parent: any = ""): any {
  function transformPrimitive(typ: string, val: any): any {
    if (typeof typ === typeof val) return val;
    return invalidValue(typ, val, key, parent);
  }

  function transformUnion(typs: any[], val: any): any {
    // val must validate against one typ in typs
    const l = typs.length;
    for (let i = 0; i < l; i++) {
      const typ = typs[i];
      try {
        return transform(val, typ, getProps);
      } catch (_) {}
    }
    return invalidValue(typs, val, key, parent);
  }

  function transformEnum(cases: string[], val: any): any {
    if (cases.indexOf(val) !== -1) return val;
    return invalidValue(
      cases.map((a) => {
        return createLiteral(a);
      }),
      val,
      key,
      parent
    );
  }

  function transformArray(typ: any, val: any): any {
    // val must be an array with no invalid elements
    if (!Array.isArray(val)) return invalidValue(createLiteral("array"), val, key, parent);
    return val.map((el) => transform(el, typ, getProps));
  }

  function transformDate(val: any): any {
    if (val === null) {
      return null;
    }
    const d = new Date(val);
    if (isNaN(d.valueOf())) {
      return invalidValue(createLiteral("Date"), val, key, parent);
    }
    return d;
  }

  function transformObject(props: { [k: string]: any }, additional: any, val: any): any {
    if (val === null || typeof val !== "object" || Array.isArray(val)) {
      return invalidValue(createLiteral(ref || "object"), val, key, parent);
    }
    const result: any = {};
    Object.getOwnPropertyNames(props).forEach((key) => {
      const prop = props[key];
      const v = Object.prototype.hasOwnProperty.call(val, key) ? val[key] : undefined;
      result[prop.key] = transform(v, prop.typ, getProps, key, ref);
    });
    Object.getOwnPropertyNames(val).forEach((key) => {
      if (!Object.prototype.hasOwnProperty.call(props, key)) {
        result[key] = transform(val[key], additional, getProps, key, ref);
      }
    });
    return result;
  }

  if (typ === "any") return val;
  if (typ === null) {
    if (val === null) return val;
    return invalidValue(typ, val, key, parent);
  }
  if (typ === false) return invalidValue(typ, val, key, parent);
  let ref: any = undefined;
  while (typeof typ === "object" && typ.ref !== undefined) {
    ref = typ.ref;
    typ = typeMap[typ.ref];
  }
  if (Array.isArray(typ)) return transformEnum(typ, val);
  if (typeof typ === "object") {
    return typ.hasOwnProperty("unionMembers")
      ? transformUnion(typ.unionMembers, val)
      : typ.hasOwnProperty("arrayItems")
        ? transformArray(typ.arrayItems, val)
        : typ.hasOwnProperty("props")
          ? transformObject(getProps(typ), typ.additional, val)
          : invalidValue(typ, val, key, parent);
  }
  // Numbers can be parsed by Date but shouldn't be.
  if (typ === Date && typeof val !== "number") return transformDate(val);
  return transformPrimitive(typ, val);
}

function cast<T>(val: any, typ: any): T {
  return transform(val, typ, jsonToJSProps);
}

function uncast<T>(val: T, typ: any): any {
  return transform(val, typ, jsToJSONProps);
}

function createLiteral(typ: any) {
  return { literal: typ };
}

function createArrayItems(typ: any) {
  return { arrayItems: typ };
}

function createUnionMembers(...typs: any[]) {
  return { unionMembers: typs };
}

function createPropsWithAdditional(props: any[], additional: any) {
  return { props, additional };
}

function createPropsArrayWithAdditional(additional: any) {
  return { props: [], additional };
}

function createReference(name: string) {
  return { ref: name };
}
