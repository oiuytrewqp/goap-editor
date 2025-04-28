import React from 'react';
import { NavLink } from 'react-router-dom';
import { Users, Brain, Target, Play, Map, FlaskConical, Menu, X } from 'lucide-react';

interface SidebarProps {
  isMobileMenuOpen: boolean;
  toggleMobileMenu: () => void;
}

const Sidebar: React.FC<SidebarProps> = ({ isMobileMenuOpen, toggleMobileMenu }) => {
  const navItems = [
    { name: 'Agents', icon: <Users size={22} />, path: '/agents' },
    { name: 'Beliefs', icon: <Brain size={22} />, path: '/beliefs' },
    { name: 'Goals', icon: <Target size={22} />, path: '/goals' },
    { name: 'Actions', icon: <Play size={22} />, path: '/actions' },
    { name: 'Locations', icon: <Map size={22} />, path: '/locations' },
    { name: 'Test', icon: <FlaskConical size={22} />, path: '/test' },
  ];

  return (
    <>
      {/* Mobile menu button */}
      <button
        onClick={toggleMobileMenu}
        className="md:hidden fixed top-4 left-4 z-20 bg-indigo-600 p-2 rounded-md"
        aria-label="Toggle menu"
      >
        {isMobileMenuOpen ? <X size={24} /> : <Menu size={24} />}
      </button>

      {/* Sidebar for mobile (slide-in) */}
      <aside 
        className={`fixed inset-y-0 left-0 w-64 bg-gray-800 transform transition-transform duration-300 ease-in-out z-10 md:hidden ${
          isMobileMenuOpen ? 'translate-x-0' : '-translate-x-full'
        }`}
      >
        <div className="h-full flex flex-col">
          <div className="p-5 flex items-center justify-center border-b border-gray-700">
            <h1 className="text-xl font-bold text-indigo-400">GOAP System</h1>
          </div>
          <nav className="flex-1 px-2 py-4 space-y-1">
            {navItems.map((item) => (
              <NavLink
                key={item.path}
                to={item.path}
                onClick={toggleMobileMenu}
                className={({ isActive }) =>
                  `flex items-center px-4 py-3 text-sm font-medium rounded-md transition-colors ${
                    isActive
                      ? 'bg-indigo-600 text-white'
                      : 'text-gray-300 hover:bg-gray-700 hover:text-white'
                  }`
                }
              >
                <span className="mr-3">{item.icon}</span>
                {item.name}
              </NavLink>
            ))}
          </nav>
        </div>
      </aside>

      {/* Sidebar for desktop (always visible) */}
      <aside className="hidden md:block w-64 bg-gray-800 overflow-y-auto">
        <div className="h-full flex flex-col">
          <div className="p-5 flex items-center justify-center border-b border-gray-700">
            <h1 className="text-xl font-bold text-indigo-400">GOAP System</h1>
          </div>
          <nav className="flex-1 px-2 py-4 space-y-1">
            {navItems.map((item) => (
              <NavLink
                key={item.path}
                to={item.path}
                className={({ isActive }) =>
                  `flex items-center px-4 py-3 text-sm font-medium rounded-md transition-colors ${
                    isActive
                      ? 'bg-indigo-600 text-white'
                      : 'text-gray-300 hover:bg-gray-700 hover:text-white'
                  }`
                }
              >
                <span className="mr-3">{item.icon}</span>
                {item.name}
              </NavLink>
            ))}
          </nav>
        </div>
      </aside>
    </>
  );
};

export default Sidebar;