import React, { useState } from "react";
import JobForm from "./JobForm";
import jobsubmitter from "./../util/jobsubmitter.js";
import FormStatus from "./FormStatus";

function NewJob(props) {
  const [status, setStatus] = useState();
  const pending = "pending";
  const success = "success";
  const result = "ok";
  const getParams = (location) => {
    const searchParams = new URLSearchParams(location.search)
    return {
      done: searchParams.get('r') || '',
    }
  };

  const onSubmit = ({ contact_name, company, email, description, title, salary, employment_type, currency_type, website, phone, city}) => {
    setStatus({ type: pending });

    jobsubmitter.submit({ contact_name, company, email, description, title, salary, employment_type, currency_type, website, phone, city }).then(() => {
      setStatus({
        type: success,
      });
      window.location.href = '/jobs/new?r=' + result
    });
  };

  return (
    <>
    {getParams(window.location).done === result && (
      <FormStatus type={success} message={props.message} />
    )}
    <JobForm
      parentColor={props.parentColor}
      buttonText={props.buttonText}
      onSubmit={onSubmit}
      status={status}
    />
    </>
  );
}

export default NewJob;
