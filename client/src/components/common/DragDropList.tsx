import React from 'react';
import { DragDropContext, Droppable, Draggable } from 'react-beautiful-dnd';
import { GripVertical, X } from 'lucide-react';

interface Item {
  id: string;
  name: string;
}

interface DragDropListProps {
  label: string;
  items: Item[];
  onReorder: (items: Item[]) => void;
  onRemove?: (id: string) => void;
}

const DragDropList: React.FC<DragDropListProps> = ({
  label,
  items,
  onReorder,
  onRemove,
}) => {
  const handleDragEnd = (result: any) => {
    if (!result.destination) {
      return;
    }

    const reorderedItems = Array.from(items);
    const [removed] = reorderedItems.splice(result.source.index, 1);
    reorderedItems.splice(result.destination.index, 0, removed);

    onReorder(reorderedItems);
  };

  return (
    <div className="mb-4">
      <label className="block text-sm font-medium text-gray-200 mb-1">{label}</label>
      <div className="bg-gray-700 rounded-md p-2">
        {items.length === 0 ? (
          <div className="text-gray-400 p-2">No items added</div>
        ) : (
          <DragDropContext onDragEnd={handleDragEnd}>
            <Droppable droppableId="droppable-list">
              {(provided) => (
                <ul
                  {...provided.droppableProps}
                  ref={provided.innerRef}
                  className="space-y-2"
                >
                  {items.map((item, index) => (
                    <Draggable key={item.id} draggableId={item.id} index={index}>
                      {(provided) => (
                        <li
                          ref={provided.innerRef}
                          {...provided.draggableProps}
                          className="bg-gray-600 rounded-md p-2 flex items-center"
                        >
                          <div
                            {...provided.dragHandleProps}
                            className="mr-2 cursor-move text-gray-400"
                          >
                            <GripVertical size={18} />
                          </div>
                          <span className="flex-1">{item.name}</span>
                          {onRemove && (
                            <button
                              onClick={() => onRemove(item.id)}
                              className="text-gray-400 hover:text-red-400"
                            >
                              <X size={18} />
                            </button>
                          )}
                        </li>
                      )}
                    </Draggable>
                  ))}
                  {provided.placeholder}
                </ul>
              )}
            </Droppable>
          </DragDropContext>
        )}
      </div>
    </div>
  );
};

export default DragDropList;