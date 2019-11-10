import React, { useState } from "react";
import FormStatus from "./FormStatus";
import FormField from "./FormField";
import SelectField from "./SelectField";
import SectionButton from "./SectionButton";
import Editor from "./Editor";

function JobForm(props) {
  const [currency_type, setCurrencyType] = useState({ value: 'rub', label: 'RUB' });
  const [employment_type, setEmploymentType] = useState({ value: 'fulltime', label: 'Fulltime' });
  const [company, setCompany] = useState("");
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [city, setCity] = useState("");
  const [phone, setPhone] = useState("");
  const [contact_name, setContactName] = useState("");
  const [email, setEmail] = useState("");
  const [salary, setSalary] = useState("");
  const [website, setWebsite] = useState("");
  const [showErrors, setShowErrors] = useState(false);

  let errors = [];

  const getError = field => {
    return errors.find(e => e.field === field);
  };

  const isEmpty = val => val.trim() === "";

  if (isEmpty(email) ) {
    errors.push({
      field: "email",
      message: "Введите email"
    });
  }

  if (isEmpty(title) ) {
    errors.push({
      field: "title",
      message: "Введите заголовок вакансии"
    });
  }

  if (isEmpty(website) ) {
    errors.push({
      field: "website",
      message: "Введите сайт компании"
    });
  }

  if (isEmpty(city) ) {
    errors.push({
      field: "city",
      message: "Введите город"
    });
  }

  if (isEmpty(company) ) {
    errors.push({
      field: "company",
      message: "Введите компанию"
    });
  }

  if (isEmpty(company) ) {
    errors.push({
      field: "company",
      message: "Введите компанию"
    });
  }

  if (isEmpty(description) ) {
    errors.push({
      field: "description",
      message: "Введите текст вакансии"
    });
  }

  const handleSubmit = e => {
    if (errors.length) {
      setShowErrors(true);
    } else {
      if (props.onSubmit) {
        props.onSubmit({
          title,
          company,
          website,
          currency_type: currency_type.value,
          employment_type: employment_type.value,
          city,
          description,
          salary: parseInt(salary),
          phone,
          contact_name,
          email
        });
        setTitle('')
        setCompany('')
        setWebsite('')
        setCity('')
        setSalary('')
        setPhone('')
        setContactName('')
        setEmail('')
        setDescription('')
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
          handleSubmit(e);
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
              placeholder="Наименование компании"
              error={showErrors && getError("company")}
              onChange={value => setCompany(value)}
            />
          </div>
        </div>

        <div className="field is-horizontal">
          <div className="field-body">
            <FormField
              value={website}
              type="text"
              placeholder="Сайт компании"
              error={showErrors && getError("website")}
              onChange={value => setWebsite(value)}
            />
          </div>
        </div>

        <div className="field is-horizontal">
          <div className="field-body">
            <FormField
              value={salary}
              type="text"
              placeholder="Заработная плата"
              error={showErrors && getError("salary")}
              onChange={value => setSalary(value)}
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
                value={contact_name}
                type="text"
                placeholder="Контактное лицо"
                error={showErrors && getError("contact_name")}
                onChange={value => setContactName(value)}
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
              value={phone}
              type="text"
              placeholder="Телефон"
              error={showErrors && getError("phone")}
              onChange={value => setPhone(value)}
            />
          </div>
        </div>

        <div className="field is-horizontal">
          <div className="field-body">
            <Editor
              value={description}
              type="textarea"
              options={{
                autofocus: true,
                spellChecker: false
              }}
              error={showErrors && getError("description")}
              onChange={value => setDescription(value)}
            />
          </div>
        </div>
        
        <div className="field is-horizontal">
          <div className="field-body">
            <div className="field">
              <div className="control">
                <SectionButton
                  parentColor={props.parentColor}
                  size="large"
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
