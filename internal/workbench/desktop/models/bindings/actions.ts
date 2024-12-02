// This file was generated by [ts-rs](https://github.com/Aleph-Alpha/ts-rs). Do not edit this file manually.
//
// The necessary import statements have been automatically added by a Python script.
// This ensures that all required dependencies are correctly referenced and available
// within this module.
//
// If you need to add or modify imports, please update the import_map in the script and
// re-run `make gen-models` it to regenerate the file accordingly.

import type { LocalizedString } from "@repo/moss-text";

export type Action = string;

export type ActionMenuItem = {
  command: CommandAction;
  group?: MenuGroup;
  order?: bigint;
  when?: object;
  visibility: MenuItemVisibility;
};

export type CommandAction = {
  id: string;
  title: LocalizedString;
  tooltip?: string;
  description?: LocalizedString;
  icon?: string;
  toggled?: CommandActionToggle;
};

export type CommandActionToggle = { condition: string; icon?: string; tooltip?: string; title?: LocalizedString };

export type MenuGroup = { id: string; order?: bigint; description?: LocalizedString };

export type MenuItem = { action: ActionMenuItem } | { submenu: SubmenuMenuItem };

export type MenuItemVisibility = "classic" | "hidden" | "compact";

export type SubmenuMenuItem = {
  submenuId: string;
  defaultActionId?: string;
  title?: LocalizedString;
  group?: MenuGroup;
  order?: bigint;
  when?: string;
  visibility: MenuItemVisibility;
};
