import React from "react";
import JobSection from "./../components/JobSection";

function JobPage(props) {
  return (
    <JobSection
      id={props.match.params.id}
      color="white"
      size="medium"
    />
  );
}

export default JobPage;
