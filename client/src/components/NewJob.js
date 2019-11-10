import React, { useState } from "react";
import JobForm from "./JobForm";
import jobsubmitter from "./../util/jobsubmitter.js";

function NewJob(props) {
  const [status, setStatus] = useState();

  const onSubmit = ({ contact_name, company, email, description, title, salary, employment_type, currency_type, website, phone, city}) => {
    setStatus({ type: "pending" });

    jobsubmitter.submit({ contact_name, company, email, description, title, salary, employment_type, currency_type, website, phone, city }).then(() => {
      setStatus({
        type: "success",
        message: "Спасибо, вакансия успешно добавлена и появится на сайте после модерации"
      });
    });
  };
  return (
    <JobForm
      parentColor={props.parentColor}
      buttonText={props.buttonText}
      onSubmit={onSubmit}
      status={status}
    />
  );
}

export default NewJob;
