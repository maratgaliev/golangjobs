import React, { useState } from "react";
import JobForm from "./JobForm";
import contact from "./../util/contact.js";

function NewJob(props) {
  const [status, setStatus] = useState();

  const onSubmit = ({ name, email, message }) => {
    setStatus({ type: "pending" });

    contact.submit({ name, email, message }).then(() => {
      setStatus({
        type: "success",
        message: "Your message has been sent! We'll get back to you soon."
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
