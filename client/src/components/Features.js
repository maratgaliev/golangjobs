import React from "react";
import "./Features.scss";
import dayjs from 'dayjs';

function Features(props) {
  return (
    <div className="Features">
      <div className='column is-narrow'>
      {props.items.map((item, index) => (
        <section key={index} className="hero is-light m-b-md">
          <div className="hero-body">
            <div className="container">
              <div className="columns is-variable bd-klmn-columns is-3">
                <div className="column is-8">
                  <h3 className="title is-uppercase">
                    {item.title}
                  </h3>
                  <h4 className="subtitle is-uppercase has-text-grey">
                    {item.company} - {item.city}
                  </h4>
                </div>
                <div className="column is-4 has-text-right">
                  <p className=''>{dayjs(item.created_at).format('YYYY-MM-DD')}</p>
                  <span className="tag is-medium is-info">УДАЛЕНКА</span>&nbsp;
                  <span className="tag is-medium is-info">FULLTIME</span>
                </div>
              </div>
            </div>
          </div>
        </section>
      ))}
      </div>
    </div>
  );
}

export default Features;
