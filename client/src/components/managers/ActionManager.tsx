import React, { useState } from 'react';
import { useAppContext } from '../../context/AppContext';
import EntityList from '../common/EntityList';
import ActionForm from '../forms/ActionForm';

const ActionManager: React.FC = () => {
  const { actions = [], selectedItem, selectItem, createItem, updateItem, deleteItem } = useAppContext();
  const [isCreating, setIsCreating] = useState(false);

  const handleSelectAction = (id: string) => {
    selectItem('action', id);
    setIsCreating(false);
  };

  const handleAddNew = () => {
    selectItem(null, null);
    setIsCreating(true);
  };

  const handleSave = async (action: any) => {
    try {
      if (isCreating) {
        await createItem('action', action);
        setIsCreating(false);
      } else if (selectedItem.id) {
        await updateItem('action', selectedItem.id, action);
      }
    } catch (error) {
      console.error('Failed to save action:', error);
    }
  };

  const handleDelete = async () => {
    if (selectedItem.id && confirm('Are you sure you want to delete this action?')) {
      await deleteItem('action', selectedItem.id);
    }
  };

  const handleCancel = () => {
    if (isCreating) {
      setIsCreating(false);
    } else {
      selectItem(null, null);
    }
  };

  const selectedAction = selectedItem.type === 'action' && selectedItem.id
    ? actions.find(action => action.id === selectedItem.id)
    : null;

  return (
    <div className="h-full">
      <h1 className="text-2xl font-bold mb-6">Action Manager</h1>
      <div className="flex h-[calc(100%-4rem)] gap-6">
        {/* Left panel - List of actions */}
        <div className="w-1/3">
          <EntityList
            title="Actions"
            items={Array.isArray(actions) ? actions.map(a => ({ id: a.id, name: a.name })) : []}
            selectedId={selectedItem.id}
            onSelect={handleSelectAction}
            onAddNew={handleAddNew}
          />
        </div>
        
        {/* Right panel - Action details */}
        <div className="w-2/3 bg-gray-800 rounded-lg overflow-hidden">
          {isCreating ? (
            <ActionForm
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={undefined}
              isNew={true}
            />
          ) : selectedAction ? (
            <ActionForm
              action={selectedAction}
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={handleDelete}
              isNew={false}
            />
          ) : (
            <div className="flex items-center justify-center h-full text-gray-500">
              <p>Select an action or create a new one</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default ActionManager;