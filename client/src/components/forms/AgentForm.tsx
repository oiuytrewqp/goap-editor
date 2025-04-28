import React, { useState, useEffect } from 'react';
import { useAppContext } from '../../context/AppContext';
import { Agent } from '../../types';
import MultiSelect from '../common/MultiSelect';
import DragDropList from '../common/DragDropList';

interface AgentFormProps {
  agent?: Agent;
  onSave: (agent: Omit<Agent, 'id'>) => void;
  onCancel: () => void;
  onDelete?: () => void;
  isNew: boolean;
}

const AgentForm: React.FC<AgentFormProps> = ({
  agent,
  onSave,
  onCancel,
  onDelete,
  isNew,
}) => {
  const { beliefs, goals, actions, locations } = useAppContext();
  
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [selectedLocationId, setSelectedLocationId] = useState<string>('');
  const [selectedBeliefIds, setSelectedBeliefIds] = useState<string[]>([]);
  const [selectedGoalIds, setSelectedGoalIds] = useState<string[]>([]);
  const [selectedActionIds, setSelectedActionIds] = useState<string[]>([]);
  
  useEffect(() => {
    if (agent) {
      setName(agent.name);
      setDescription(agent.description);
      setSelectedLocationId(agent.locationId);
      setSelectedBeliefIds(agent.beliefs);
      setSelectedGoalIds(agent.goals);
      setSelectedActionIds(agent.actions);
    } else {
      // Default values for new agent
      setName('');
      setDescription('');
      setSelectedLocationId(locations.length > 0 ? locations[0].id : '');
      setSelectedBeliefIds([]);
      setSelectedGoalIds([]);
      setSelectedActionIds([]);
    }
  }, [agent, locations]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    const updatedAgent = {
      name,
      description,
      locationId: selectedLocationId,
      beliefs: selectedBeliefIds,
      goals: selectedGoalIds,
      actions: selectedActionIds,
    };
    
    onSave(updatedAgent);
  };

  const handleGoalReorder = (reorderedGoals: { id: string; name: string }[]) => {
    setSelectedGoalIds(reorderedGoals.map(goal => goal.id));
  };

  const handleRemoveGoal = (id: string) => {
    setSelectedGoalIds(selectedGoalIds.filter(goalId => goalId !== id));
  };

  return (
    <form onSubmit={handleSubmit} className="h-full flex flex-col">
      <div className="p-6 border-b border-gray-700">
        <h2 className="text-xl font-semibold mb-4">
          {isNew ? 'Create New Agent' : 'Edit Agent'}
        </h2>
      </div>
      
      <div className="flex-1 overflow-y-auto p-6">
        <div className="mb-4">
          <label htmlFor="name" className="block text-sm font-medium text-gray-200 mb-1">
            Name
          </label>
          <input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            required
          />
        </div>

        <div className="mb-4">
          <label htmlFor="description" className="block text-sm font-medium text-gray-200 mb-1">
            Description
          </label>
          <textarea
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            rows={3}
            className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
          />
        </div>

        <div className="mb-4">
          <label htmlFor="location" className="block text-sm font-medium text-gray-200 mb-1">
            Location
          </label>
          <select
            id="location"
            value={selectedLocationId}
            onChange={(e) => setSelectedLocationId(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            required
          >
            <option value="" disabled>
              Select a location
            </option>
            {locations.map((location) => (
              <option key={location.id} value={location.id}>
                {location.name}
              </option>
            ))}
          </select>
        </div>

        <MultiSelect
          label="Beliefs"
          options={beliefs.map(belief => ({ id: belief.id, name: belief.key }))}
          selectedIds={selectedBeliefIds}
          onChange={setSelectedBeliefIds}
        />
        
        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-200 mb-1">
            Goals (Priority Order)
          </label>
          <DragDropList
            label="Goals (Priority Order)"
            items={selectedGoalIds.map(id => {
              const goal = goals.find(g => g.id === id);
              return { id, name: goal ? goal.name : 'Unknown Goal' };
            })}
            onReorder={handleGoalReorder}
            onRemove={handleRemoveGoal}
          />
          
          <div className="mt-2">
            <select
              className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
              value=""
              onChange={(e) => {
                if (e.target.value && !selectedGoalIds.includes(e.target.value)) {
                  setSelectedGoalIds([...selectedGoalIds, e.target.value]);
                }
                e.target.value = '';
              }}
            >
              <option value="" disabled>
                Add a goal...
              </option>
              {goals
                .filter(goal => !selectedGoalIds.includes(goal.id))
                .map(goal => (
                  <option key={goal.id} value={goal.id}>
                    {goal.name}
                  </option>
                ))}
            </select>
          </div>
        </div>

        <MultiSelect
          label="Actions"
          options={actions.map(action => ({ id: action.id, name: action.name }))}
          selectedIds={selectedActionIds}
          onChange={setSelectedActionIds}
        />
      </div>
      
      <div className="p-6 border-t border-gray-700 flex justify-between">
        <div>
          {!isNew && onDelete && (
            <button
              type="button"
              onClick={onDelete}
              className="px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-md mr-2 transition-colors"
            >
              Delete
            </button>
          )}
        </div>
        <div>
          <button
            type="button"
            onClick={onCancel}
            className="px-4 py-2 bg-gray-600 hover:bg-gray-700 text-white rounded-md mr-2 transition-colors"
          >
            Cancel
          </button>
          <button
            type="submit"
            className="px-4 py-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded-md transition-colors"
          >
            Save
          </button>
        </div>
      </div>
    </form>
  );
};

export default AgentForm;