import React from "react";
import Navbar from "./../components/Navbar";
import IndexPage from "./index";
import AboutPage from "./about";
import FaqPage from "./faq";
import ContactPage from "./contact";
import { Switch, Route, Router } from "./../util/router.js";
import Footer from "./../components/Footer";
import "./../util/analytics.js";
import { ProvideAuth } from "./../util/auth.js";

function App(props) {
  return (
    <ProvideAuth>
      <Router>
        <>
          <Navbar
            color="white"
            spaced={true}
            logo="https://golang.org/lib/godoc/images/go-logo-blue.svg"
          />

          <Switch>
            <Route exact path="/" component={IndexPage} />

            <Route exact path="/about" component={AboutPage} />

            <Route exact path="/faq" component={FaqPage} />

            <Route exact path="/contact" component={ContactPage} />

            <Route
              component={({ location }) => {
                return (
                  <div
                    style={{
                      padding: "50px",
                      width: "100%",
                      textAlign: "center"
                    }}
                  >
                    Страница <code>{location.pathname}</code> не найдена.
                  </div>
                );
              }}
            />
          </Switch>

          <Footer
            color="light"
            size="normal"
            logo="https://golang.org/lib/godoc/images/go-logo-blue.svg"
            copyright="© 2019 GOLANGJOBS"
          />
        </>
      </Router>
    </ProvideAuth>
  );
}

export default App;
