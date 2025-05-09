import React, { useState, useEffect } from 'react';
import { useAppContext } from '../../context/AppContext';
import { Action } from '../../types';
import MultiSelect from '../common/MultiSelect';

interface ActionFormProps {
  action?: Action;
  onSave: (action: Omit<Action, 'id'>) => void;
  onCancel: () => void;
  onDelete?: () => void;
  isNew: boolean;
}

const ActionForm: React.FC<ActionFormProps> = ({
  action,
  onSave,
  onCancel,
  onDelete,
  isNew,
}) => {
  const { beliefs, locations } = useAppContext();
  
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [methodName, setMethodName] = useState('');
  const [selectedLocationId, setSelectedLocationId] = useState<string | null>(null);
  const [prerequisiteIds, setPrerequisiteIds] = useState<string[]>([]);
  const [outcomeIds, setOutcomeIds] = useState<string[]>([]);
  
  useEffect(() => {
    if (action) {
      setName(action.name);
      setDescription(action.description);
      setMethodName(action.methodName);
      setSelectedLocationId(action.locationId);
      setPrerequisiteIds(action.prerequisites);
      setOutcomeIds(action.outcomes);
    } else {
      // Default values for new action
      setName('');
      setDescription('');
      setMethodName('');
      setSelectedLocationId(null);
      setPrerequisiteIds([]);
      setOutcomeIds([]);
    }
  }, [action]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    const updatedAction = {
      name,
      description,
      methodName,
      locationId: selectedLocationId,
      prerequisites: prerequisiteIds,
      outcomes: outcomeIds,
    };
    
    onSave(updatedAction);
  };

  return (
    <form onSubmit={handleSubmit} className="h-full flex flex-col">
      <div className="p-6 border-b border-gray-700">
        <h2 className="text-xl font-semibold mb-4">
          {isNew ? 'Create New Action' : 'Edit Action'}
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
          <label htmlFor="methodName" className="block text-sm font-medium text-gray-200 mb-1">
            Method Name
          </label>
          <input
            type="text"
            id="methodName"
            value={methodName}
            onChange={(e) => setMethodName(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            placeholder="e.g., moveToLocation, pickUpItem"
          />
        </div>

        <div className="mb-4">
          <label htmlFor="location" className="block text-sm font-medium text-gray-200 mb-1">
            Location (Optional)
          </label>
          <select
            id="location"
            value={selectedLocationId || ''}
            onChange={(e) => setSelectedLocationId(e.target.value || null)}
            className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
          >
            <option value="">None (No location required)</option>
            {locations.map((location) => (
              <option key={location.id} value={location.id}>
                {location.name}
              </option>
            ))}
          </select>
        </div>

        <MultiSelect
          label="Prerequisites (Beliefs required to perform this action)"
          options={beliefs.map(belief => ({ id: belief.id, name: belief.name }))}
          selectedIds={prerequisiteIds}
          onChange={setPrerequisiteIds}
        />

        <MultiSelect
          label="Outcomes (Beliefs resulting from this action)"
          options={beliefs.map(belief => ({ id: belief.id, name: belief.name }))}
          selectedIds={outcomeIds}
          onChange={setOutcomeIds}
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

export default ActionForm;