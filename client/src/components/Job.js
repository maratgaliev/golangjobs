import React from "react";
import JobBody from "./JobBody";

function Job(props) {
  return (
    <JobBody
      id={props.id}
      job={props.job}
      parentColor={props.parentColor}
      buttonText={props.buttonText}
    />
  );
}

export default Job;
