import { useTranslation } from "react-i18next";
import { commands, SessionInfoDTO } from "@/bindings";
import React, { useEffect, useState } from "react";
import { listen } from "@tauri-apps/api/event";
import {
  Tooltip,
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  Icon,
} from "@repo/ui";
import { invokeIpc } from "@/lib/backend/tauri";

export type DescribeActivityOutput = { tooltip: string; order: number };

const SessionComponent = () => {
  const { t } = useTranslation(["ns1", "ns2"]);
  const [sessionInfo, setSessionInfo] = useState<SessionInfoDTO | null>(null);
  const [data, setData] = useState<number | null>(null);
  const [workbenchState, setWorkbenchState] = useState<string>("empty");

  let getWorkbenchState = async () => {
    try {
      const response = await commands.workbenchGetState();
      if (response.status === "ok") {
        setWorkbenchState(response.data);
      }
    } catch (err) {
      console.error("Failed to get workbench state:", err);
    }
  };

  let getAllActivities = async () => {
    try {
      // console.log((await invokeIpc("describe_primary_activitybar_part")) as object);
      // console.log((await invokeIpc("describe_primary_sidebar_part")) as object);
      // console.log((await invokeIpc("get_view_content")) as object);
      // console.log((await invokeCmd("get_menu_items")) as object);
    } catch (err) {
      console.error("Failed to get workbench state:", err);
    }
  };

  useEffect(() => {
    const unlisten = listen<number>("data-stream", (event) => {
      setData(event.payload);
    });

    getWorkbenchState();
    getAllActivities();

    return () => {
      unlisten.then((f) => f());
    };
  }, []);

  useEffect(() => {
    if (sessionInfo) {
      console.log("Session restored:", sessionInfo);
    }
  }, [sessionInfo]);

  const handleRestoreSession = async () => {
    try {
      let response = await commands.restoreSession(null);
      if (response.status === "ok") {
        setSessionInfo(response.data);
      }
    } catch (error) {
      console.error("Failed to restore session:", error);
    }
  };

  return (
    <>
      {sessionInfo ? (
        <div>
          <p>Session: {sessionInfo.session.id}</p>
          <p>Project: {sessionInfo.project.source}</p>
        </div>
      ) : (
        <p>No session</p>
      )}

      <p className="text-[rgba(var(--color-primary))]">
        Workspace: <span className="bg-red-500 text-[rgba(var(--color-primary))]"> {workbenchState}</span>
      </p>
      <br />

      <button className="bg-red-500 text-[rgba(var(--color-primary))]" onClick={handleRestoreSession}>
        Restore Session
      </button>
      <br />

      <span className="text-[rgba(var(--color-primary))]">{t("description.part1")}</span>
      <br />
      <span className="bg-secondary text-[rgba(var(--color-primary))]">{t("description.part1", { ns: "ns2" })}</span>
      {data !== null && <p>Received data: {data}</p>}
    </>
  );
};

export const Home: React.FC = () => {
  const { t } = useTranslation(["ns1", "ns2"]);

  const handleNewWindowButton = async () => {
    const response = await invokeIpc("create_new_window");
    console.log(response);
  };

  return (
    <div>
      <h1 className="text-[rgba(var(--color-primary))]">{t("title")}</h1>

      <button className="bg-green-500 px-3" onClick={handleNewWindowButton}>
        New Window
      </button>

      <div>
        <Tooltip label="Test" className="text-[rgba(var(--color-primary))]">
          <Icon icon="Code" />
        </Tooltip>
      </div>
      <SessionComponent />
      <div>
        {/* <DropdownMenu>
          <DropdownMenuTrigger className="text-[rgba(var(--color-primary))]">Click me!</DropdownMenuTrigger>

          <DropdownMenuContent>
            <DropdownMenuItem icon="Search">Menu item 1</DropdownMenuItem>
            <DropdownMenuItem>
              <DropdownMenuLabel>Menu item 2</DropdownMenuLabel>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu> */}
      </div>

      <div className="flex">
        <Icon icon="Accessibility" className="text-6xl hover:*:fill-green-500" />
        <Icon icon="NewProject" className="text-red-700 text-6xl hover:fill-green-500" />
      </div>

      <div className="w-96 bg-red-600">
        {new Array(77).fill(0).map((_, index) => (
          <div key={index}>
            {index} - {Math.random()}
          </div>
        ))}
        <div> last element</div>
      </div>
    </div>
  );
};
