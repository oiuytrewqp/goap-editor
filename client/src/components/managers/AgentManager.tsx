import React, { useState } from 'react';
import { useAppContext } from '../../context/AppContext';
import EntityList from '../common/EntityList';
import AgentForm from '../forms/AgentForm';

const AgentManager: React.FC = () => {
  const { agents, selectedItem, selectItem, createItem, updateItem, deleteItem } = useAppContext();
  const [isCreating, setIsCreating] = useState(false);

  const handleSelectAgent = (id: string) => {
    selectItem('agent', id);
    setIsCreating(false);
  };

  const handleAddNew = () => {
    selectItem(null, null);
    setIsCreating(true);
  };

  const handleSave = async (agent: any) => {
    try {
      if (isCreating) {
        await createItem('agent', agent);
        setIsCreating(false);
      } else if (selectedItem.id) {
        await updateItem('agent', selectedItem.id, agent);
      }
    } catch (error) {
      console.error('Failed to save agent:', error);
    }
  };

  const handleDelete = async () => {
    if (selectedItem.id && confirm('Are you sure you want to delete this agent?')) {
      await deleteItem('agent', selectedItem.id);
    }
  };

  const handleCancel = () => {
    if (isCreating) {
      setIsCreating(false);
    } else {
      selectItem(null, null);
    }
  };

  const selectedAgent = selectedItem.type === 'agent' && selectedItem.id
    ? agents.find(agent => agent.id === selectedItem.id)
    : null;

  const agentList = Array.isArray(agents) ? agents : [];

  return (
    <div className="h-full">
      <h1 className="text-2xl font-bold mb-6">Agent Manager</h1>
      <div className="flex h-[calc(100%-4rem)] gap-6">
        {/* Left panel - List of agents */}
        <div className="w-1/3">
          <EntityList
            title="Agents"
            items={agentList.map(a => ({ id: a.id, name: a.name }))}
            selectedId={selectedItem.id}
            onSelect={handleSelectAgent}
            onAddNew={handleAddNew}
          />
        </div>
        
        {/* Right panel - Agent details */}
        <div className="w-2/3 bg-gray-800 rounded-lg overflow-hidden">
          {isCreating ? (
            <AgentForm
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={undefined}
              isNew={true}
            />
          ) : selectedAgent ? (
            <AgentForm
              agent={selectedAgent}
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={handleDelete}
              isNew={false}
            />
          ) : (
            <div className="flex items-center justify-center h-full text-gray-500">
              <p>Select an agent or create a new one</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default AgentManager;