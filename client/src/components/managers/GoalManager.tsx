import React, { useState } from 'react';
import { useAppContext } from '../../context/AppContext';
import EntityList from '../common/EntityList';
import GoalForm from '../forms/GoalForm';

const GoalManager: React.FC = () => {
  const { goals, selectedItem, selectItem, createItem, updateItem, deleteItem } = useAppContext();
  const [isCreating, setIsCreating] = useState(false);

  const handleSelectGoal = (id: string) => {
    selectItem('goal', id);
    setIsCreating(false);
  };

  const handleAddNew = () => {
    selectItem(null, null);
    setIsCreating(true);
  };

  const handleSave = async (goal: any) => {
    try {
      if (isCreating) {
        await createItem('goal', goal);
        setIsCreating(false);
      } else if (selectedItem.id) {
        await updateItem('goal', selectedItem.id, goal);
      }
    } catch (error) {
      console.error('Failed to save goal:', error);
    }
  };

  const handleDelete = async () => {
    if (selectedItem.id && confirm('Are you sure you want to delete this goal?')) {
      await deleteItem('goal', selectedItem.id);
    }
  };

  const handleCancel = () => {
    if (isCreating) {
      setIsCreating(false);
    } else {
      selectItem(null, null);
    }
  };

  const selectedGoal = selectedItem.type === 'goal' && selectedItem.id
    ? (Array.isArray(goals) ? goals.find(goal => goal.id === selectedItem.id) : null)
    : null;

  return (
    <div className="h-full">
      <h1 className="text-2xl font-bold mb-6">Goal Manager</h1>
      <div className="flex h-[calc(100%-4rem)] gap-6">
        {/* Left panel - List of goals */}
        <div className="w-1/3">
          <EntityList
            title="Goals"
            items={Array.isArray(goals) ? goals.map(g => ({ id: g.id, name: g.name })) : []}
            selectedId={selectedItem.id}
            onSelect={handleSelectGoal}
            onAddNew={handleAddNew}
          />
        </div>
        
        {/* Right panel - Goal details */}
        <div className="w-2/3 bg-gray-800 rounded-lg overflow-hidden">
          {isCreating ? (
            <GoalForm
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={undefined}
              isNew={true}
            />
          ) : selectedGoal ? (
            <GoalForm
              goal={selectedGoal}
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={handleDelete}
              isNew={false}
            />
          ) : (
            <div className="flex items-center justify-center h-full text-gray-500">
              <p>Select a goal or create a new one</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default GoalManager;