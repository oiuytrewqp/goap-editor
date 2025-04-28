import React, { useState, useEffect } from 'react';
import { Location } from '../../types';

interface LocationFormProps {
  location?: Location;
  onSave: (location: Omit<Location, 'id'>) => void;
  onCancel: () => void;
  onDelete?: () => void;
  isNew: boolean;
}

const LocationForm: React.FC<LocationFormProps> = ({
  location,
  onSave,
  onCancel,
  onDelete,
  isNew,
}) => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  
  useEffect(() => {
    if (location) {
      setName(location.name);
      setDescription(location.description);
    } else {
      // Default values for new location
      setName('');
      setDescription('');
    }
  }, [location]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    const updatedLocation = {
      name,
      description,
    };
    
    onSave(updatedLocation);
  };

  return (
    <form onSubmit={handleSubmit} className="h-full flex flex-col">
      <div className="p-6 border-b border-gray-700">
        <h2 className="text-xl font-semibold mb-4">
          {isNew ? 'Create New Location' : 'Edit Location'}
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

export default LocationForm;