import React, { useState } from 'react';
import { Search, Plus } from 'lucide-react';

interface EntityListProps {
  title: string;
  items: Array<{ id: string; name: string }>;
  selectedId: string | null;
  onSelect: (id: string) => void;
  onAddNew: () => void;
}

const EntityList: React.FC<EntityListProps> = ({
  title,
  items,
  selectedId,
  onSelect,
  onAddNew,
}) => {
  const [searchTerm, setSearchTerm] = useState('');

  const filteredItems = items.filter((item) =>
    item.name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="flex flex-col h-full bg-gray-800 rounded-lg overflow-hidden">
      <div className="p-4 bg-gray-700">
        <h2 className="text-lg font-semibold mb-3">{title}</h2>
        <div className="relative">
          <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
            <Search size={18} className="text-gray-400" />
          </div>
          <input
            type="text"
            className="bg-gray-600 text-white placeholder-gray-400 w-full pl-10 pr-4 py-2 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
            placeholder="Search..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
      </div>
      
      <div className="flex-1 overflow-y-auto p-2">
        {filteredItems.length > 0 ? (
          <ul className="space-y-1">
            {filteredItems.map((item) => (
              <li key={item.id}>
                <button
                  onClick={() => onSelect(item.id)}
                  className={`w-full text-left px-4 py-3 rounded-md transition-colors ${
                    selectedId === item.id
                      ? 'bg-indigo-600 text-white'
                      : 'hover:bg-gray-700 text-gray-300'
                  }`}
                >
                  {item.name}
                </button>
              </li>
            ))}
          </ul>
        ) : (
          <div className="text-center py-6 text-gray-500">
            {searchTerm ? 'No results found' : 'No items yet'}
          </div>
        )}
      </div>

      <div className="p-4 border-t border-gray-700">
        <button
          onClick={onAddNew}
          className="w-full flex items-center justify-center px-4 py-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded-md transition-colors"
        >
          <Plus size={18} className="mr-2" />
          Add New
        </button>
      </div>
    </div>
  );
};

export default EntityList;