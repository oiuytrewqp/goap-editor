import React, { useState, useEffect } from 'react';
import { Belief } from '../../types';

interface BeliefFormProps {
  belief?: Belief;
  onSave: (belief: Omit<Belief, 'id'>) => void;
  onCancel: () => void;
  onDelete?: () => void;
  isNew: boolean;
}

const BeliefForm: React.FC<BeliefFormProps> = ({
  belief,
  onSave,
  onCancel,
  onDelete,
  isNew,
}) => {
  const [key, setKey] = useState('');
  const [value, setValue] = useState<any>('');
  const [valueType, setValueType] = useState<'string' | 'number' | 'boolean'>('string');
  const [description, setDescription] = useState('');
  
  useEffect(() => {
    if (belief) {
      setKey(belief.key);
      
      // Determine value type and set accordingly
      if (typeof belief.value === 'boolean') {
        setValueType('boolean');
        setValue(belief.value);
      } else if (typeof belief.value === 'number') {
        setValueType('number');
        setValue(belief.value);
      } else {
        setValueType('string');
        setValue(String(belief.value));
      }
      
      setDescription(belief.description);
    } else {
      // Default values for new belief
      setKey('');
      setValue('');
      setValueType('string');
      setDescription('');
    }
  }, [belief]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    
    // Convert value to the correct type
    let typedValue;
    if (valueType === 'boolean') {
      typedValue = value === 'true';
    } else if (valueType === 'number') {
      typedValue = Number(value);
    } else {
      typedValue = value;
    }
    
    const updatedBelief = {
      key,
      value: typedValue,
      description,
    };
    
    onSave(updatedBelief);
  };

  const handleValueTypeChange = (newType: 'string' | 'number' | 'boolean') => {
    setValueType(newType);
    
    // Reset value when changing types
    if (newType === 'boolean') {
      setValue('false');
    } else if (newType === 'number') {
      setValue('0');
    } else {
      setValue('');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="h-full flex flex-col">
      <div className="p-6 border-b border-gray-700">
        <h2 className="text-xl font-semibold mb-4">
          {isNew ? 'Create New Belief' : 'Edit Belief'}
        </h2>
      </div>
      
      <div className="flex-1 overflow-y-auto p-6">
        <div className="mb-4">
          <label htmlFor="key" className="block text-sm font-medium text-gray-200 mb-1">
            Key
          </label>
          <input
            type="text"
            id="key"
            value={key}
            onChange={(e) => setKey(e.target.value)}
            className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            required
          />
        </div>

        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-200 mb-1">
            Value Type
          </label>
          <div className="flex gap-4">
            <label className="inline-flex items-center">
              <input
                type="radio"
                className="text-indigo-600"
                checked={valueType === 'string'}
                onChange={() => handleValueTypeChange('string')}
              />
              <span className="ml-2">String</span>
            </label>
            <label className="inline-flex items-center">
              <input
                type="radio"
                className="text-indigo-600"
                checked={valueType === 'number'}
                onChange={() => handleValueTypeChange('number')}
              />
              <span className="ml-2">Number</span>
            </label>
            <label className="inline-flex items-center">
              <input
                type="radio"
                className="text-indigo-600"
                checked={valueType === 'boolean'}
                onChange={() => handleValueTypeChange('boolean')}
              />
              <span className="ml-2">Boolean</span>
            </label>
          </div>
        </div>

        <div className="mb-4">
          <label htmlFor="value" className="block text-sm font-medium text-gray-200 mb-1">
            Value
          </label>
          {valueType === 'boolean' ? (
            <select
              id="value"
              value={value}
              onChange={(e) => setValue(e.target.value)}
              className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            >
              <option value="true">True</option>
              <option value="false">False</option>
            </select>
          ) : (
            <input
              type={valueType === 'number' ? 'number' : 'text'}
              id="value"
              value={value}
              onChange={(e) => setValue(e.target.value)}
              className="w-full bg-gray-700 text-white rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            />
          )}
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

export default BeliefForm;