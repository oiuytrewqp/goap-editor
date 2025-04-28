import React, { useState } from 'react';
import { useAppContext } from '../../context/AppContext';
import EntityList from '../common/EntityList';
import BeliefForm from '../forms/BeliefForm';

const BeliefManager: React.FC = () => {
  const { beliefs = [], selectedItem, selectItem, createItem, updateItem, deleteItem } = useAppContext();
  const [isCreating, setIsCreating] = useState(false);

  const handleSelectBelief = (id: string) => {
    selectItem('belief', id);
    setIsCreating(false);
  };

  const handleAddNew = () => {
    selectItem(null, null);
    setIsCreating(true);
  };

  const handleSave = async (belief: any) => {
    try {
      if (isCreating) {
        await createItem('belief', belief);
        setIsCreating(false);
      } else if (selectedItem.id) {
        await updateItem('belief', selectedItem.id, belief);
      }
    } catch (error) {
      console.error('Failed to save belief:', error);
    }
  };

  const handleDelete = async () => {
    if (selectedItem.id && confirm('Are you sure you want to delete this belief?')) {
      await deleteItem('belief', selectedItem.id);
    }
  };

  const handleCancel = () => {
    if (isCreating) {
      setIsCreating(false);
    } else {
      selectItem(null, null);
    }
  };

  const selectedBelief = selectedItem.type === 'belief' && selectedItem.id && Array.isArray(beliefs)
    ? beliefs.find(belief => belief.id === selectedItem.id)
    : null;

  const beliefItems = Array.isArray(beliefs) 
    ? beliefs.map(b => ({ id: b.id, name: b.key }))
    : [];

  return (
    <div className="h-full">
      <h1 className="text-2xl font-bold mb-6">Belief Manager</h1>
      <div className="flex h-[calc(100%-4rem)] gap-6">
        {/* Left panel - List of beliefs */}
        <div className="w-1/3">
          <EntityList
            title="Beliefs"
            items={beliefItems}
            selectedId={selectedItem.id}
            onSelect={handleSelectBelief}
            onAddNew={handleAddNew}
          />
        </div>
        
        {/* Right panel - Belief details */}
        <div className="w-2/3 bg-gray-800 rounded-lg overflow-hidden">
          {isCreating ? (
            <BeliefForm
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={undefined}
              isNew={true}
            />
          ) : selectedBelief ? (
            <BeliefForm
              belief={selectedBelief}
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={handleDelete}
              isNew={false}
            />
          ) : (
            <div className="flex items-center justify-center h-full text-gray-500">
              <p>Select a belief or create a new one</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default BeliefManager;