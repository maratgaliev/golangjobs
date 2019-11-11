import React from "react";
import Section from "./Section";
import SectionHeader from "./SectionHeader";
import NewJob from "./NewJob";
import "./NewJobSection.scss";

function NewJobSection(props) {
  return (
    <Section color={props.color} size={props.size}>
      <div className="NewJobSection__container container">
        <SectionHeader
          title={props.title}
          subtitle={props.subtitle}
          centered={true}
          size={3}
        />
        <NewJob
          message={props.message}
          parentColor={props.color}
          buttonText={props.buttonText}
        />
      </div>
    </Section>
  );
}

export default NewJobSection;
