import React, { useState } from "react";
import FormStatus from "./FormStatus";
import FormField from "./FormField";
import SelectField from "./SelectField";
import SectionButton from "./SectionButton";
import SimpleMDE from "react-simplemde-editor";
import "easymde/dist/easymde.min.css";

function JobForm(props) {
  // State for input values
  const [currency_type, setCurrencyType] = useState({ value: 'rub', label: 'RUB' });
  const [employment_type, setEmploymentType] = useState({ value: 'fulltime', label: 'Fulltime' });
  const [company, setCompany] = useState("");
  const [title, setTitle] = useState("");
  const [city, setCity] = useState("");
  const [contact_phone, setContactPhone] = useState("");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [message, setMessage] = useState("");

  // Whether to show errors
  // We set to true if they submit and there are errors.
  // We only show errors after they submit because
  // it's annoying to see errors while typing.
  const [showErrors, setShowErrors] = useState(false);

  // Error array we'll populate
  let errors = [];

  // Function for fetching error for a field
  const getError = field => {
    return errors.find(e => e.field === field);
  };

  // Function to see if field is empty
  const isEmpty = val => val.trim() === "";

  // Add error if email empty
  if (isEmpty(email) ) {
    errors.push({
      field: "email",
      message: "Please enter value"
    });
  }

  // Handle form submission
  const handleSubmit = e => {
    // If field errors then show them
    if (errors.length) {
      setShowErrors(true);
    } else {
      // Otherwise call onSubmit with form data
      if (props.onSubmit) {
        props.onSubmit({
          title,
          company,
          currency_type,
          employment_type,
          city,
          contact_phone,
          name,
          email,
          message
        });
      }
    }
  };

  return (
    <>
      {props.status && props.status.message && (
        <FormStatus type={props.status.type} message={props.status.message} />
      )}

      <form
        onSubmit={e => {
          e.preventDefault();
          handleSubmit();
        }}
      >
        <div className="field is-horizontal">
          <div className="field-body">
            <FormField
              value={title}
              type="text"
              placeholder="Заголовок вакансии"
              error={showErrors && getError("title")}
              onChange={value => setTitle(value)}
            />
          </div>
        </div>

        <div className="field is-horizontal">
          <div className="field-body">
            <FormField
              value={company}
              type="text"
              placeholder="Наименование организации"
              error={showErrors && getError("company")}
              onChange={value => setCompany(value)}
            />
          </div>
        </div>

        <div className="field is-horizontal">
          <div className="field-body">
              <SelectField
                value={employment_type}
                options={[
                  { value: 'fulltime', label: 'Fulltime' },
                  { value: 'remote', label: 'Remote' },
                  { value: 'contract', label: 'Contract' },
                  { value: 'freelance', label: 'Freelance' },
                ]}
                placeholder="Тип занятости"
                error={showErrors && getError("employment_type")}
                onChange={value => setEmploymentType(value)}
              />
              <SelectField
                value={currency_type}
                options={[
                  { value: 'eur', label: 'EUR' },
                  { value: 'usd', label: 'USD' },
                  { value: 'rub', label: 'RUB' }
                ]}
                placeholder="Валюта оплаты"
                error={showErrors && getError("currency_type")}
                onChange={value => setCurrencyType(value)}
              />
          </div>
        </div>

        <div className="field is-horizontal">
          <div className="field-body">
              <FormField
                value={name}
                type="text"
                placeholder="Контактное лицо"
                error={showErrors && getError("name")}
                onChange={value => setName(value)}
              />
            <FormField
              value={email}
              type="email"
              placeholder="Email"
              error={showErrors && getError("email")}
              onChange={value => setEmail(value)}
            />
          </div>
        </div>

        <div className="field is-horizontal">
          <div className="field-body">
              <FormField
                value={city}
                type="text"
                placeholder="Город"
                error={showErrors && getError("city")}
                onChange={value => setCity(value)}
              />
            <FormField
              value={contact_phone}
              type="text"
              placeholder="Телефон"
              error={showErrors && getError("contactPhone")}
              onChange={value => setContactPhone(value)}
            />
          </div>
        </div>

        <div className="field is-horizontal">
          <div className="field-body">
            <SimpleMDE
              value={message}
              type="textarea"
              placeholder="Текст вакансии"
              options={{
                autofocus: true,
                spellChecker: false
              }}
              //error={showErrors && getError("message")}
              onChange={value => setMessage(value)}
            />
          </div>
        </div>
        <div className="field is-horizontal">
          <div className="field-body">
            <div className="field">
              <div className="control">
                <SectionButton
                  parentColor={props.parentColor}
                  size="medium"
                  state={
                    props.status && props.status.type === "pending"
                      ? "loading"
                      : "normal"
                  }
                >
                  {props.buttonText}
                </SectionButton>
              </div>
            </div>
          </div>
        </div>
      </form>
    </>
  );
}

export default JobForm;
