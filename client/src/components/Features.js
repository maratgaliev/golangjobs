import React from "react";
import "./Features.scss";

function Features(props) {
  return (
    <div className="Features">
      <div className='column is-narrow'>
      {props.items.map((item, index) => (
        <article key={index} className="message is-info">
        <div className="message-header">
          <p>{item.title}</p>
        </div>
        <div className="message-body">
          {item.description}
        </div>
      </article>
      ))}
      </div>
    </div>
  );
}

export default Features;
