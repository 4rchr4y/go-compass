import {
  DndContext,
  DragEndEvent,
  DragOverlay,
  DragStartEvent,
  KeyboardSensor,
  MouseSensor,
  UniqueIdentifier,
  closestCenter,
  useSensor,
  useSensors,
} from "@dnd-kit/core";
import {
  SortableContext,
  arrayMove,
  horizontalListSortingStrategy,
  sortableKeyboardCoordinates,
} from "@dnd-kit/sortable";
import { Icon, cn } from "@repo/ui";
import { OsType } from "@tauri-apps/plugin-os";
import { HTMLProps, useState } from "react";
import { createPortal } from "react-dom";
import { HeadBarButton } from "./HeadBarButton";

interface WidgetBarProps extends HTMLProps<HTMLDivElement> {
  os: OsType;
}

export const WidgetBar = ({ os, className, ...props }: WidgetBarProps) => {
  const [activeId, setActiveId] = useState<UniqueIdentifier | null>(null);

  const [items, setItems] = useState([
    {
      id: 1,
      label: "Alerts",
      icon: "HeadBarAlerts" as const,
    },
    {
      id: 2,
      label: "Discovery",
      icon: "HeadBarDiscovery" as const,
    },
    {
      id: 3,
      label: "Community",
      icon: "HeadBarCommunity" as const,
    },
  ]);

  function handleDragStart(event: DragStartEvent) {
    setActiveId(event.active.id);
  }

  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;

    setActiveId(null);

    if (!over) return;

    if (active.id !== over.id) {
      setItems((items) => {
        const oldIndex = items.findIndex((a) => a.id === active.id);
        const newIndex = items.findIndex((a) => a.id === over.id);

        return arrayMove(items, oldIndex, newIndex);
      });
    }
  };

  const sensors = useSensors(
    useSensor(MouseSensor, {
      activationConstraint: {
        distance: 5,
      },
    }),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  );

  return (
    <div className={cn("flex items-center gap-1", className)} {...props}>
      {os !== "macos" && (
        <HeadBarButton
          icon="HeadBarSettingsWithNotification"
          className="flex size-[30px] items-center justify-center px-2"
          iconClassName="size-[18px]"
        />
      )}
      <div className="flex items-center gap-3">
        <button className="flex h-[30px] w-max items-center rounded pl-2.5 pr-1 transition-colors hover:bg-[#D3D3D3]">
          <Icon icon="HeadBarMossStudio" className="mr-1.5 size-[22px] text-[#525252]" />
          <span className="mr-0.5 w-max text-[#161616]">moss-studio</span>
          <Icon icon="ArrowheadDown" className="text-[#525252]" />
        </button>

        <div className="flex w-full justify-between">
          <div className={cn("flex items-center gap-1")}>
            <DndContext
              sensors={sensors}
              collisionDetection={closestCenter}
              onDragEnd={handleDragEnd}
              onDragStart={handleDragStart}
            >
              <SortableContext items={items} strategy={horizontalListSortingStrategy}>
                {items.map((item) => (
                  <HeadBarButton
                    key={item.id}
                    sortableId={item.id}
                    icon={item.icon}
                    label={item.label}
                    className={cn("h-[30px] text-ellipsis px-2")}
                  />
                ))}
              </SortableContext>

              {activeId
                ? createPortal(
                    <DragOverlay>
                      <HeadBarButton
                        className="h-[30px] cursor-grabbing !bg-[#e0e0e0] px-2 shadow-lg"
                        icon={items.find((item) => item.id === activeId)?.icon!}
                        label={items.find((item) => item.id === activeId)?.label}
                      />
                    </DragOverlay>,
                    document.body
                  )
                : null}
            </DndContext>
          </div>
        </div>
      </div>
    </div>
  );
};
