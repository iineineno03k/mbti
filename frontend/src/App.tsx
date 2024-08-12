import React from 'react';
import './App.css';
import UserTable from './component/UserTable';
import UserFormModal from './component/UserFormModal';

function App() {
  return (
    <div>
      <UserFormModal />
      <UserTable />
    </div>
  );
}

export default App;
