import React from "react";
import SimpleMDE from "react-simplemde-editor";
import "easymde/dist/easymde.min.css";

function Editor(props) {
  return (
    <div className="field">
      <div className="control">
        <SimpleMDE
          id='custom_id'
          value={props.value}
          options={{
            autofocus: true,
            spellChecker: false
          }}
          onChange={props.onChange}
        />
      </div>

      {props.error && <p className="help is-danger">{props.error.message}</p>}
    </div>
  );
}

export default Editor;
