import React from "react";
import Section from "./Section";
import SectionHeader from "./SectionHeader";
import Clients from "./Clients";

function ClientsSection(props) {
  return (
    <Section color={props.color} size={props.size}>
      <div className="container">
        <SectionHeader
          title={props.title}
          subtitle={props.subtitle}
          centered={true}
          size={3}
        />
        <Clients
          items={[
            {
              name: "Instagram",
              image: "logo-instagram.svg",
              width: "150px"
            },
            {
              name: "Instagram",
              image: "logo-instagram.svg",
              width: "150px"
            },
            {
              name: "Instagram",
              image: "logo-instagram.svg",
              width: "150px"
            },
            {
              name: "Instagram",
              image: "logo-instagram.svg",
              width: "150px"
            }
          ]}
        />
      </div>
    </Section>
  );
}

export default ClientsSection;
