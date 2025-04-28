import React, { useState, useEffect } from 'react';
import { useAppContext } from '../../context/AppContext';
import { Agent, Belief, Location } from '../../types';
import { Play, RefreshCw } from 'lucide-react';

const TestEnvironment: React.FC = () => {
  const { 
    agents, 
    beliefs, 
    locations,
    loading, 
    stepSimulation,
    resetSimulation
  } = useAppContext();
  
  const [stepCount, setStepCount] = useState(0);
  const [logMessages, setLogMessages] = useState<string[]>([]);

  // Group agents by location
  const agentsByLocation = locations.reduce<Record<string, Agent[]>>((acc, location) => {
    acc[location.id] = agents.filter(agent => agent.locationId === location.id);
    return acc;
  }, {});

  const handleStep = async () => {
    try {
      await stepSimulation();
      setStepCount(prevCount => prevCount + 1);
      setLogMessages(prev => [`Step ${stepCount + 1} completed`, ...prev.slice(0, 19)]);
    } catch (error) {
      console.error('Error stepping simulation:', error);
      setLogMessages(prev => [`Error: Failed to step simulation`, ...prev.slice(0, 19)]);
    }
  };

  const handleReset = async () => {
    try {
      await resetSimulation();
      setStepCount(0);
      setLogMessages(['Simulation reset']);
    } catch (error) {
      console.error('Error resetting simulation:', error);
      setLogMessages(prev => [`Error: Failed to reset simulation`, ...prev.slice(0, 19)]);
    }
  };

  // Find beliefs for an agent
  const getAgentBeliefs = (agentId: string): Belief[] => {
    const agent = agents.find(a => a.id === agentId);
    if (!agent) return [];
    
    return beliefs.filter(belief => agent.beliefs.includes(belief.id));
  };

  return (
    <div className="h-full">
      <h1 className="text-2xl font-bold mb-6">Test Environment</h1>
      
      <div className="mb-4 flex justify-between items-center">
        <div>
          <span className="text-gray-300 mr-4">Current Step: {stepCount}</span>
          <button
            onClick={handleStep}
            disabled={loading}
            className="px-4 py-2 bg-indigo-600 hover:bg-indigo-700 text-white rounded-md mr-2 transition-colors flex items-center"
          >
            <Play size={18} className="mr-2" />
            Step
          </button>
          <button
            onClick={handleReset}
            disabled={loading}
            className="px-4 py-2 bg-gray-600 hover:bg-gray-700 text-white rounded-md transition-colors flex items-center"
          >
            <RefreshCw size={18} className="mr-2" />
            Reset
          </button>
        </div>
        
        {loading && (
          <div className="text-yellow-400 animate-pulse">Processing...</div>
        )}
      </div>
      
      <div className="flex gap-6 h-[calc(100%-12rem)]">
        {/* Left side - Locations and agents */}
        <div className="w-2/3 bg-gray-800 rounded-lg overflow-hidden">
          <div className="p-4 bg-gray-700 border-b border-gray-600">
            <h2 className="text-lg font-semibold">World State</h2>
          </div>
          
          <div className="p-4 overflow-y-auto h-[calc(100%-60px)]">
            {locations.length > 0 ? (
              <div className="grid grid-cols-1 gap-6">
                {locations.map((location) => (
                  <div key={location.id} className="bg-gray-700 rounded-lg overflow-hidden">
                    <div className="bg-gray-600 p-3 border-b border-gray-500">
                      <h3 className="font-medium">{location.name}</h3>
                      {location.description && (
                        <p className="text-sm text-gray-400 mt-1">{location.description}</p>
                      )}
                    </div>
                    
                    <div className="p-4">
                      <h4 className="text-sm text-gray-400 mb-2">Agents at this location:</h4>
                      {agentsByLocation[location.id]?.length > 0 ? (
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-3">
                          {agentsByLocation[location.id].map((agent) => (
                            <div key={agent.id} className="bg-gray-800 p-3 rounded-md border border-gray-700">
                              <div className="font-medium text-indigo-400 mb-1">{agent.name}</div>
                              <div className="text-xs text-gray-400 mb-2">{agent.description}</div>
                              
                              {getAgentBeliefs(agent.id).length > 0 && (
                                <div>
                                  <div className="text-xs text-gray-400 mb-1">Current beliefs:</div>
                                  <div className="flex flex-wrap gap-2">
                                    {getAgentBeliefs(agent.id).map((belief) => (
                                      <span 
                                        key={belief.id} 
                                        className="text-xs bg-gray-700 text-gray-300 px-2 py-1 rounded"
                                      >
                                        {belief.key}: {String(belief.value)}
                                      </span>
                                    ))}
                                  </div>
                                </div>
                              )}
                            </div>
                          ))}
                        </div>
                      ) : (
                        <p className="text-sm text-gray-500 italic">No agents at this location</p>
                      )}
                    </div>
                  </div>
                ))}
              </div>
            ) : (
              <div className="flex items-center justify-center h-full text-gray-500">
                <p>No locations defined yet. Create some locations first.</p>
              </div>
            )}
          </div>
        </div>
        
        {/* Right side - Log */}
        <div className="w-1/3 bg-gray-800 rounded-lg overflow-hidden">
          <div className="p-4 bg-gray-700 border-b border-gray-600">
            <h2 className="text-lg font-semibold">Simulation Log</h2>
          </div>
          
          <div className="p-4 overflow-y-auto h-[calc(100%-60px)]">
            {logMessages.length > 0 ? (
              <div className="space-y-2">
                {logMessages.map((message, index) => (
                  <div key={index} className="text-sm">
                    <span className="text-gray-400 mr-2">[{new Date().toLocaleTimeString()}]</span>
                    <span>{message}</span>
                  </div>
                ))}
              </div>
            ) : (
              <p className="text-gray-500 text-center py-4">
                No activity yet. Click "Step" to start the simulation.
              </p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default TestEnvironment;