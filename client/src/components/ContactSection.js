import React from "react";
import Section from "./Section";
import "./ContactSection.scss";

function ContactSection(props) {
  return (
    <Section color={props.color} size={props.size}>
      <div className="container">
        <div className="content is-medium has-text-left">
          <p>По всем идеям и предложениям:</p> 
          <p>Сайт: <a target='_blank' href="http://maratgaliev.com">http://maratgaliev.com</a></p>
          <p>Почта: <a href='mailto:kazanlug@gmail.com'>kazanlug@gmail.com</a></p>
        </div>
      </div>
    </Section>
  );
}

export default ContactSection;
