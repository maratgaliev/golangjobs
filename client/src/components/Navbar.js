import React, { useState } from "react";
import NavbarContainer from "./NavbarContainer";
import { Link } from "./../util/router.js";

function Navbar(props) {
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <NavbarContainer spaced={props.spaced} color={props.color}>
      <div className="container">
        <div className="navbar-brand">
          <div className="navbar-item">
            <Link to="/">
              <img className="image" src={process.env.PUBLIC_URL + '/' + props.logo} alt="Logo" />
            </Link>
          </div>
          <div
            className={"navbar-burger burger" + (menuOpen ? " is-active" : "")}
            onClick={() => setMenuOpen(!menuOpen)}
          >
            <span />
            <span />
            <span />
          </div>
        </div>
        <div className={"navbar-menu" + (menuOpen ? " is-active" : "")}>
          <div className="navbar-start">
            <Link className="navbar-item" to="/">
              ГЛАВНАЯ
            </Link>
            <Link className="navbar-item" to="/events">
              СОБЫТИЯ
            </Link>
            <Link className="navbar-item" to="/about">
              О ПРОЕКТЕ
            </Link>
            <Link className="navbar-item" to="/contact">
              КОНТАКТЫ
            </Link>
            <Link className="navbar-item" to="/faq">
              FAQ
            </Link>
          </div>

          <div className="navbar-end">
            <div className="navbar-item">
              <div className="buttons">
                <Link className="navbar-item button is-info" to="/jobs/new">
                  ДОБАВИТЬ ВАКАНСИЮ
                </Link>
              </div>
            </div>
          </div>
        </div>
      </div>
    </NavbarContainer>
  );
}

export default Navbar;
