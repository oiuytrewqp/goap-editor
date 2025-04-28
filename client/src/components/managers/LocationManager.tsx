import React, { useState } from 'react';
import { useAppContext } from '../../context/AppContext';
import EntityList from '../common/EntityList';
import LocationForm from '../forms/LocationForm';

const LocationManager: React.FC = () => {
  const { locations, selectedItem, selectItem, createItem, updateItem, deleteItem } = useAppContext();
  const [isCreating, setIsCreating] = useState(false);

  const handleSelectLocation = (id: string) => {
    selectItem('location', id);
    setIsCreating(false);
  };

  const handleAddNew = () => {
    selectItem(null, null);
    setIsCreating(true);
  };

  const handleSave = async (location: any) => {
    try {
      if (isCreating) {
        await createItem('location', location);
        setIsCreating(false);
      } else if (selectedItem.id) {
        await updateItem('location', selectedItem.id, location);
      }
    } catch (error) {
      console.error('Failed to save location:', error);
    }
  };

  const handleDelete = async () => {
    if (selectedItem.id && confirm('Are you sure you want to delete this location?')) {
      await deleteItem('location', selectedItem.id);
    }
  };

  const handleCancel = () => {
    if (isCreating) {
      setIsCreating(false);
    } else {
      selectItem(null, null);
    }
  };

  const selectedLocation = selectedItem.type === 'location' && selectedItem.id
    ? (Array.isArray(locations) ? locations.find(location => location.id === selectedItem.id) : null)
    : null;

  return (
    <div className="h-full">
      <h1 className="text-2xl font-bold mb-6">Location Manager</h1>
      <div className="flex h-[calc(100%-4rem)] gap-6">
        {/* Left panel - List of locations */}
        <div className="w-1/3">
          <EntityList
            title="Locations"
            items={Array.isArray(locations) ? locations.map(l => ({ id: l.id, name: l.name })) : []}
            selectedId={selectedItem.id}
            onSelect={handleSelectLocation}
            onAddNew={handleAddNew}
          />
        </div>
        
        {/* Right panel - Location details */}
        <div className="w-2/3 bg-gray-800 rounded-lg overflow-hidden">
          {isCreating ? (
            <LocationForm
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={undefined}
              isNew={true}
            />
          ) : selectedLocation ? (
            <LocationForm
              location={selectedLocation}
              onSave={handleSave}
              onCancel={handleCancel}
              onDelete={handleDelete}
              isNew={false}
            />
          ) : (
            <div className="flex items-center justify-center h-full text-gray-500">
              <p>Select a location or create a new one</p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};

export default LocationManager;