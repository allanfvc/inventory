import React from 'react'
import './header.css'

const Header: React.FC = (props) => {
  return (
    <header>
      <nav className="navbar navbar-dark bg-primary">
      <div className="container-fluid">
        Inventário
      </div>
      </nav>
    </header>
  );
}

export default Header
