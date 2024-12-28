import * as React from "react";

import {
  DockviewApi,
  DockviewDefaultTab,
  DockviewReact,
  DockviewReadyEvent,
  IDockviewPanelHeaderProps,
  IDockviewPanelProps,
} from "@repo/moss-tabs";

import "./app.scss";

import { LeftControls, PrefixHeaderControls, RightControls } from "./controls";
import { Table, usePanelApiMetadata } from "./debugPanel";
import { defaultConfig } from "./defaultLayout";
import { GridActions } from "./gridActions";
import { GroupActions } from "./groupActions";
import { PanelActions } from "./panelActions";

const DebugContext = React.createContext<boolean>(false);

const Option = (props: { title: string; onClick: () => void; value: string }) => {
  return (
    <div>
      <span>{`${props.title}: `}</span>
      <button onClick={props.onClick}>{props.value}</button>
    </div>
  );
};

const components = {
  Default: (props: IDockviewPanelProps) => {
    const isDebug = React.useContext(DebugContext);
    const metadata = usePanelApiMetadata(props.api);

    return (
      <div
        className={`relative h-full overflow-auto p-1.5 ${isDebug ? "border-2 border-dashed border-orange-500" : ""}`}
      >
        <span className="pointer-events-none absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2 transform text-6xl opacity-50">
          {props.api.title}
        </span>

        {isDebug && (
          <div className="text-xs">
            <Option
              title="Panel Rendering Mode"
              value={metadata.renderer.value}
              onClick={() => props.api.setRenderer(props.api.renderer === "always" ? "onlyWhenVisible" : "always")}
            />

            <Table data={metadata} />
          </div>
        )}
      </div>
    );
  },
  nested: (props: IDockviewPanelProps) => {
    return (
      <DockviewReact
        components={components}
        onReady={(event: DockviewReadyEvent) => {
          event.api.addPanel({ id: "panel_1", component: "Default" });
          event.api.addPanel({ id: "panel_2", component: "Default" });
          event.api.addPanel({
            id: "panel_3",
            component: "Default",
            floating: true,
          });
        }}
        className={"dockview-theme-abyss"}
      />
    );
  },
  iframe: (props: IDockviewPanelProps) => {
    return (
      <iframe
        onMouseDown={() => {
          if (!props.api.isActive) {
            props.api.setActive();
          }
        }}
        className="h-full w-full"
        src="https://dockview.dev"
      />
    );
  },
};

const headerComponents = {
  default: (props: IDockviewPanelHeaderProps) => {
    const onContextMenu = (event: React.MouseEvent) => {
      event.preventDefault();
      alert("context menu");
    };
    return <DockviewDefaultTab onContextMenu={onContextMenu} {...props} />;
  },
};

const colors = [
  "rgba(255,0,0,0.2)",
  "rgba(0,255,0,0.2)",
  "rgba(0,0,255,0.2)",
  "rgba(255,255,0,0.2)",
  "rgba(0,255,255,0.2)",
  "rgba(255,0,255,0.2)",
];
let count = 0;

const WatermarkComponent = () => {
  return <div>custom watermark</div>;
};

const DockviewDemo = (props: { theme?: string }) => {
  const [logLines, setLogLines] = React.useState<{ text: string; timestamp?: Date; backgroundColor?: string }[]>([]);

  const [panels, setPanels] = React.useState<string[]>([]);
  const [groups, setGroups] = React.useState<string[]>([]);
  const [api, setApi] = React.useState<DockviewApi>();

  const [activePanel, setActivePanel] = React.useState<string>();
  const [activeGroup, setActiveGroup] = React.useState<string>();

  const [pending, setPending] = React.useState<{ text: string; timestamp?: Date }[]>([]);

  const addLogLine = (message: string) => {
    setPending((line) => [{ text: message, timestamp: new Date() }, ...line]);
  };

  React.useLayoutEffect(() => {
    if (pending.length === 0) {
      return;
    }
    const color = colors[count++ % colors.length];
    setLogLines((lines) => [...pending.map((_) => ({ ..._, backgroundColor: color })), ...lines]);
    setPending([]);
  }, [pending]);

  const onReady = (event: DockviewReadyEvent) => {
    setApi(event.api);

    event.api.onDidAddPanel((event) => {
      setPanels((_) => [..._, event.id]);
      addLogLine(`Panel Added ${event.id}`);
    });
    event.api.onDidActivePanelChange((event) => {
      setActivePanel(event?.id);
      addLogLine(`Panel Activated ${event?.id}`);
    });
    event.api.onDidRemovePanel((event) => {
      setPanels((_) => {
        const next = [..._];
        next.splice(
          next.findIndex((x) => x === event.id),
          1
        );

        return next;
      });
      addLogLine(`Panel Removed ${event.id}`);
    });

    event.api.onDidAddGroup((event) => {
      setGroups((_) => [..._, event.id]);
      addLogLine(`Group Added ${event.id}`);
    });

    event.api.onDidMovePanel((event) => {
      addLogLine(`Panel Moved ${event.panel.id}`);
    });

    event.api.onDidMaximizedGroupChange((event) => {
      addLogLine(`Group Maximized Changed ${event.view.id} [${event.isMaximized}]`);
    });

    event.api.onDidRemoveGroup((event) => {
      setGroups((_) => {
        const next = [..._];
        next.splice(
          next.findIndex((x) => x === event.id),
          1
        );

        return next;
      });
      addLogLine(`Group Removed ${event.id}`);
    });

    event.api.onDidActiveGroupChange((event) => {
      setActiveGroup(event?.id);
      addLogLine(`Group Activated ${event?.id}`);
    });

    const state = localStorage.getItem("dv-demo-state");
    if (state) {
      try {
        event.api.fromJSON(JSON.parse(state));
        return;
      } catch {
        localStorage.removeItem("dv-demo-state");
      }
      return;
    }

    defaultConfig(event.api);
  };

  const [watermark, setWatermark] = React.useState<boolean>(false);

  const [gapCheck, setGapCheck] = React.useState<boolean>(false);

  const css = React.useMemo(() => {
    if (!gapCheck) {
      return {};
    }

    return {
      "--dv-group-gap-size": "0.5rem",
      "--demo-border": "5px dashed purple",
    } as React.CSSProperties;
  }, [gapCheck]);

  const [showLogs, setShowLogs] = React.useState<boolean>(false);
  const [debug, setDebug] = React.useState<boolean>(false);

  return (
    <div className={`relative flex h-full flex-grow flex-col rounded-lg bg-[rgba(0,0,50,0.25)] p-2 ${css}`}>
      <div>
        <GridActions api={api} toggleCustomWatermark={() => setWatermark(!watermark)} hasCustomWatermark={watermark} />
        {api && <PanelActions api={api} panels={panels} activePanel={activePanel} />}
        {api && <GroupActions api={api} groups={groups} activeGroup={activeGroup} />}
        {/* <div>
                  <button
                      onClick={() => {
                          setGapCheck(!gapCheck);
                      }}
                  >
                      {gapCheck ? 'Disable Gap Check' : 'Enable Gap Check'}
                  </button>
              </div> */}
      </div>
      <div className="action-container flex items-center justify-end p-1">
        <button
          onClick={() => {
            setDebug(!debug);
          }}
        >
          <span className="material-symbols-outlined">engineering</span>
        </button>
        {showLogs && (
          <button
            onClick={() => {
              setLogLines([]);
            }}
          >
            <span className="material-symbols-outlined">undo</span>
          </button>
        )}
        <button
          onClick={() => {
            setShowLogs(!showLogs);
          }}
        >
          <span className="pr-1">{`${showLogs ? "Hide" : "Show"} Events Log`}</span>
          <span className="material-symbols-outlined">terminal</span>
        </button>
      </div>
      <div className="flex h-0 flex-grow">
        <div className="flex h-full flex-grow overflow-hidden">
          <DebugContext.Provider value={debug}>
            <DockviewReact
              components={components}
              defaultTabComponent={headerComponents.default}
              rightHeaderActionsComponent={RightControls}
              leftHeaderActionsComponent={LeftControls}
              prefixHeaderActionsComponent={PrefixHeaderControls}
              watermarkComponent={watermark ? WatermarkComponent : undefined}
              onReady={onReady}
              className={props.theme || "dockview-theme-abyss"}
            />
          </DebugContext.Provider>
        </div>

        {showLogs && (
          <div className="bg-black ml-2 w-[400px] flex-shrink-0 overflow-auto font-mono text-white">
            {logLines.map((line, i) => {
              return (
                <div
                  className="flex h-[30px] items-center overflow-hidden text-ellipsis whitespace-nowrap text-sm"
                  style={{ backgroundColor: line.backgroundColor }}
                  key={i}
                >
                  <span className="text-gray-500 border-gray-500 mr-1 flex h-full min-w-[20px] max-w-[20px] items-center border-r pl-1">
                    {logLines.length - i}
                  </span>
                  <span>
                    {line.timestamp && (
                      <span className="px-0.5 text-[0.7em]">{line.timestamp.toISOString().substring(11, 23)}</span>
                    )}
                    <span>{line.text}</span>
                  </span>
                </div>
              );
            })}
          </div>
        )}
      </div>
    </div>
  );
};

export default DockviewDemo;
