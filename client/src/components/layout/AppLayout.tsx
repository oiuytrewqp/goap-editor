import React, { useState } from 'react';
import { Routes, Route, Navigate } from 'react-router-dom';
import Sidebar from './Sidebar';
import AgentManager from '../managers/AgentManager';
import BeliefManager from '../managers/BeliefManager';
import GoalManager from '../managers/GoalManager';
import ActionManager from '../managers/ActionManager';
import LocationManager from '../managers/LocationManager';
import TestEnvironment from '../test/TestEnvironment';

const AppLayout: React.FC = () => {
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const toggleMobileMenu = () => {
    setIsMobileMenuOpen(!isMobileMenuOpen);
  };

  return (
    <div className="flex h-screen bg-gray-900 text-white">
      <Sidebar 
        isMobileMenuOpen={isMobileMenuOpen}
        toggleMobileMenu={toggleMobileMenu}
      />
      
      <main className="flex-1 overflow-auto p-4 md:p-6">
        <Routes>
          <Route path="/agents" element={<AgentManager />} />
          <Route path="/beliefs" element={<BeliefManager />} />
          <Route path="/goals" element={<GoalManager />} />
          <Route path="/actions" element={<ActionManager />} />
          <Route path="/locations" element={<LocationManager />} />
          <Route path="/test" element={<TestEnvironment />} />
          <Route path="*" element={<Navigate to="/agents" replace />} />
        </Routes>
      </main>
    </div>
  );
};

export default AppLayout;