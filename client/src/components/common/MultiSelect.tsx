import React, { useState } from 'react';
import { Check, X } from 'lucide-react';

interface Option {
  id: string;
  name: string;
}

interface MultiSelectProps {
  label: string;
  options: Option[];
  selectedIds: string[];
  onChange: (selectedIds: string[]) => void;
}

const MultiSelect: React.FC<MultiSelectProps> = ({
  label,
  options,
  selectedIds,
  onChange,
}) => {
  const [isOpen, setIsOpen] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');

  const filteredOptions = options.filter((option) =>
    option.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const toggleOption = (id: string) => {
    const newSelectedIds = selectedIds.includes(id)
      ? selectedIds.filter((selectedId) => selectedId !== id)
      : [...selectedIds, id];
    onChange(newSelectedIds);
  };

  const removeOption = (id: string, e: React.MouseEvent) => {
    e.stopPropagation();
    onChange(selectedIds.filter((selectedId) => selectedId !== id));
  };

  return (
    <div className="mb-4">
      <label className="block text-sm font-medium text-gray-200 mb-1">{label}</label>
      <div className="relative">
        <div
          className="bg-gray-700 rounded-md p-2 min-h-[42px] flex flex-wrap gap-2 cursor-pointer"
          onClick={() => setIsOpen(!isOpen)}
        >
          {selectedIds.length > 0 ? (
            selectedIds.map((id) => {
              const option = options.find((o) => o.id === id);
              return option ? (
                <div
                  key={id}
                  className="flex items-center bg-indigo-500 text-white px-2 py-1 rounded-md text-sm"
                >
                  {option.name}
                  <button
                    onClick={(e) => removeOption(id, e)}
                    className="ml-1 text-indigo-200 hover:text-white"
                  >
                    <X size={14} />
                  </button>
                </div>
              ) : null;
            })
          ) : (
            <div className="text-gray-400 py-1">Select options...</div>
          )}
        </div>

        {isOpen && (
          <div className="absolute z-10 mt-1 w-full bg-gray-700 rounded-md shadow-lg">
            <div className="p-2">
              <input
                type="text"
                className="w-full p-2 bg-gray-600 text-white rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
                placeholder="Search..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                onClick={(e) => e.stopPropagation()}
              />
            </div>
            <ul className="max-h-60 overflow-y-auto py-1">
              {filteredOptions.map((option) => {
                const isSelected = selectedIds.includes(option.id);
                return (
                  <li key={option.id}>
                    <button
                      type="button"
                      className={`w-full flex items-center px-4 py-2 text-sm hover:bg-gray-600 ${
                        isSelected ? 'bg-indigo-500 text-white' : 'text-gray-200'
                      }`}
                      onClick={(e) => {
                        e.stopPropagation();
                        toggleOption(option.id);
                      }}
                    >
                      <span className="flex-1 text-left">{option.name}</span>
                      {isSelected && <Check size={16} />}
                    </button>
                  </li>
                );
              })}
              {filteredOptions.length === 0 && (
                <li className="px-4 py-2 text-sm text-gray-400">No results found</li>
              )}
            </ul>
          </div>
        )}
      </div>
    </div>
  );
};

export default MultiSelect;