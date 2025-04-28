import React, { useState, useEffect } from 'react';
import { useAppContext } from '../../context/AppContext';
import { Goal } from '../../types';
import MultiSelect from '../common/MultiSelect';

interface GoalFormProps {
  goal?: Goal;
  onSave: (goal: Omit<Goal, 'id'>) => void;
  onCancel: () => void;
  onDelete?: () => void;
  isNew: boolean;
}

const GoalForm: React.FC<GoalFormProps> = ({
  goal,
  onSave,
  onCancel,
  onDelete,
  isNew,
}) => {
  const { beliefs } = useAppContext();
  
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [priority, setPriority] = useState(1);
  const [selectedBeliefIds, setSelectedBeliefIds] = useState<string[]>([]);
  
  useEffect(() => {
    if (goal) {
      setName(goal.name);
      setDescription(goal.description);
      setPriority(goal.priority);
      setSelectedBeliefIds(goal.beliefs);
    } else {
      // Default values for new goal
      setName('');
      setDescription('');
      setPriority(1);
      setSelectedBeliefIds([]);
    }
  }, [goal]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    const updatedGoal = {
      name,
      description,
      priority,
      beliefs: selectedBeliefIds,
    };
    
    onSave(updatedGoal);
  };

  return (
    <form onSubmit={handleSubmit} className="h-full flex flex-col">
      <div className="p-6 border-b border-gray-700">
        <h2 className="text-xl font-semibold mb-4">
          {isNew ? 'Create New Goal' : 'Edit Goal'}
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
          <label htmlFor="priority" className="block text-sm font-medium text-gray-200 mb-1">
            Priority (1-10, 1 is highest)
          </label>
          <input
            type="number"
            id="priority"
            value={priority}
            onChange={(e) => setPriority(Math.max(1, Math.min(10, parseInt(e.target.value))))}
            min={1}
            max={10}
            className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            required
          />
        </div>

        <MultiSelect
          label="Beliefs (required to satisfy this goal)"
          options={beliefs.map(belief => ({ id: belief.id, name: belief.key }))}
          selectedIds={selectedBeliefIds}
          onChange={setSelectedBeliefIds}
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

export default GoalForm;