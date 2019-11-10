import React from "react";
import Section from "./Section";
import SectionHeader from "./SectionHeader";
import Features from "./Features";
import useFetch from "./../util/fetcher.js";

function FeaturesSection(props) {
  const { loading, data } = useFetch(
    "http://localhost:8081/api/jobs"
  );

  return (
    <div className="App">
      {loading && <div className="loader" />}
      {data &&
        data.length > 0 &&
        <>
          <Section color={props.color} size={props.size}>
            <div className="container">
              <SectionHeader
                title={props.title}
                subtitle={props.subtitle}
                centered={true}
                size={3}
              />
              <Features
                items={data}
              />
            </div>
          </Section>
        </>
        }
    </div>
  );
}

export default FeaturesSection;