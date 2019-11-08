import React from "react";
import "./Features.scss";
import dayjs from "dayjs";

function Features(props) {
  return (
    <div className="Features">
      <div className="column is-narrow">
        {props.items.map((item, index) => (
          <section key={index} className="hero is-light m-b-md">
            <header className="card-header">
              <p className="card-header-title has-text-centered is-uppercase has-text-weight-bold is-family-code">
                {item.city}
              </p>
            </header>
            <div className="hero-body">
              <div className="container">
                <div className="has-text-centered row m-b-md">
                  <h3 className="title is-uppercase has-text-weight-semibold has-text-info">
                    {item.title}
                  </h3>
                  <h4 className="subtitle is-uppercase has-text-grey">{item.company}</h4>
                </div>
                <div className="has-text-centered row m-b-md">
                  <span className="tag is-medium is-success">УДАЛЕНКА</span>
                  &nbsp;
                  <span className="tag is-medium is-success">FULLTIME</span>
                </div>
                <div className="has-text-centered row">
                  <span className="tag is-medium is-info">ПОДРОБНОСТИ</span>
                </div>
              </div>
            </div>
            <footer className="item-footer">
              <b className="card-footer-item has-text-weight-light is-family-secondary has-text-grey-dark">
                ДОБАВЛЕНО: {dayjs(item.created_at).format("DD/MM/YYYY")}
              </b>
            </footer>
          </section>
        ))}
      </div>
    </div>
  );
}

export default Features;
