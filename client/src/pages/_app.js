import React from "react";
import Navbar from "./../components/Navbar";
import IndexPage from "./index";
import AboutPage from "./about";
import FaqPage from "./faq";
import ContactPage from "./contact";
import NewJobPage from "./newjob";
import JobPage from "./job";
import { Switch, Route, Router } from "./../util/router.js";
import Footer from "./../components/Footer";
import "./../util/analytics.js";

function App(props) {
  return (
      <Router>
        <>
          <Navbar
            color="white"
            spaced={true}
            logo="go-logo-blue.svg"
          />

          <Switch>
            <Route exact path="/" component={IndexPage} />

            <Route exact path="/about" component={AboutPage} />

            <Route exact path="/faq" component={FaqPage} />

            <Route exact path="/contact" component={ContactPage} />

            <Route path="/jobs/new" component={NewJobPage} />
            
            <Route exact path="/jobs/:id" component={JobPage} />

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
            logo="go-logo-blue.svg"
            copyright="© 2019 GOLANGJOBS"
          />
        </>
      </Router>
  );
}

export default App;
