import React from "react";
import Section from "./Section";
import SectionHeader from "./SectionHeader";
import Job from "./Job";
import "./JobSection.scss";
import useFetch from "./../util/fetcher.js";

function JobSection(props) {
  const { loading, data } = useFetch(
    "http://localhost:8081/api/jobs/" + props.id
  );

  return (
    <Section color={props.color} size={props.size}>
      <div className="JobSection__container container">
        {loading && <div className="loader" />}
        {data && (
          <>
            <SectionHeader
              title={data.title}
              subtitle={data.city}
              centered={true}
              size={3}
            />
            <Job
              job={data}
              id={props.id}
              parentColor={props.color}
              buttonText={props.buttonText}
            />
          </>
        )}
      </div>
    </Section>
  );
}

export default JobSection;
