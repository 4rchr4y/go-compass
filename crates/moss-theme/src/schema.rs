use serde_json::{json, Value};

pub const SCHEMA_THEME: Value = json!({
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "Theme.json",
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "description": "The display name of the theme. It is used to identify the theme in a human-readable format."
    },
    "slug": {
      "type": "string",
      "description": "A unique identifier for the theme, used as a slug for referencing the theme in URLs or configuration."
    },
    "type": {
      "anyOf": [
        {
          "type": "string",
          "const": "light"
        },
        {
          "type": "string",
          "const": "dark"
        }
      ],
      "description": "The category or type of the theme, e.g., 'dark', 'light'."
    },
    "colors": {
      "$ref": "#/$defs/Colors",
      "description": "An object containing the color configuration for various UI elements."
    }
  },
  "required": ["name", "slug", "type", "colors"],
  "$defs": {
    "Colors": {
      "type": "object",
      "properties": {
        "primary": {
          "$ref": "#/$defs/ColorDefinition"
        },
        "sideBar.background": {
          "$ref": "#/$defs/ColorDefinition"
        },
        "toolBar.background": {
          "$ref": "#/$defs/ColorDefinition"
        },
        "page.background": {
          "$ref": "#/$defs/ColorDefinition"
        },
        "statusBar.background": {
          "$ref": "#/$defs/ColorDefinition"
        },
        "windowsCloseButton.background": {
          "$ref": "#/$defs/ColorDefinition"
        },
        "windowControlsLinux.background": {
          "$ref": "#/$defs/ColorDefinition"
        },
        "windowControlsLinux.text": {
          "$ref": "#/$defs/ColorDefinition"
        },
        "windowControlsLinux.hoverBackground": {
          "$ref": "#/$defs/ColorDefinition"
        },
        "windowControlsLinux.activeBackground": {
          "$ref": "#/$defs/ColorDefinition"
        }
      },
      "required": [
        "primary",
        "sideBar.background",
        "toolBar.background",
        "page.background",
        "statusBar.background",
        "windowsCloseButton.background",
        "windowControlsLinux.background",
        "windowControlsLinux.text",
        "windowControlsLinux.hoverBackground",
        "windowControlsLinux.activeBackground"
      ]
    },
    "ColorDefinition": {
      "anyOf": [
        {
          "$ref": "#/$defs/SolidColor"
        },
        {
          "$ref": "#/$defs/GradientColor"
        }
      ]
    },
    "SolidColor": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "const": "solid"
        },
        "value": {
          "$ref": "#/$defs/SolidString"
        }
      },
      "required": ["type", "value"]
    },
    "GradientColor": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "const": "gradient"
        },
        "value": {
          "$ref": "#/$defs/GradientString"
        }
      },
      "required": ["type", "value"]
    },
    "SolidString": {
      "type": "string",
      "pattern": "#[0-9a-fA-F]{3,8}|rgb(a)?\\(.+\\)|hsl(a)?\\(.+\\)|[a-z]+"
    },
    "GradientString": {
      "type": "string",
      "pattern": "\\w+-gradient\\(.*?\\)\\)|\\w+-gradient\\(.*?\\)"
    }
  }
}

);
