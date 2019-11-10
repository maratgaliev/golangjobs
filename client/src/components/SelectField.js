import React from "react";
import Select from 'react-select';

function SelectField(props) {
  return (
    <div className="field">
      <div className="control">
      <Select
        value={props.value}
        options={props.options}
        onChange={props.onChange}
      />
      </div>

      {props.error && <p className="help is-danger">{props.error.message}</p>}
    </div>
  );
}

export default SelectField;
