import React from "react";
import dayjs from "dayjs";
import "dayjs/locale/ru";
dayjs.locale("ru");
const ReactMarkdown = require("react-markdown");

function JobBody(props) {
  return (
    <div className="container">
      <section className="articles">
        <div className="column is-8 is-offset-2">
          <div className="card article">
            <div className="card-content">
              <div className="media">
                <div className="media-content has-text-centered">
                  <p className="title article-title is-uppercase">
                    {props.job.company}
                  </p>
                  {props.job.salary && (
                    <>
                      <p className="title salary-title is-uppercase">
                        {props.job.salary} {props.job.currency_type}
                      </p>
                      <br />
                    </>
                  )}
                  <div className="has-text-centered row m-b-md">
                    <span className="tag is-medium is-info is-uppercase">{props.job.employment_type}</span>
                  </div>
                  <br />
                  <div className="tags has-addons level-item">
                    <span className="tag is-rounded is-info">Добавлено</span>
                    <span className="tag is-rounded">
                      {dayjs(props.job.created_at).format("DD MMMM YYYY")}
                    </span>
                  </div>
                </div>
              </div>
              <div className="content article-body">
                <ReactMarkdown source={props.job.description} />
              </div>
              <hr />
              <div className="content">
                <table className="table is-fullwidth">
                  <tbody>
                    <tr>
                      <td width="5%">
                        <i className="fa fa-bell-o"></i>
                      </td>
                      <td className="is-uppercase">Сайт компании</td>
                      <td className="level-right">
                        <a rel="noopener noreferrer" target='_blank' className="button is-uppercase is-small is-info" href={props.job.website}>
                          {props.job.website}
                        </a>
                      </td>
                    </tr>
                    <tr>
                      <td width="5%">
                        <i className="fa fa-bell-o"></i>
                      </td>
                      <td className="is-uppercase">Контактное лицо</td>
                      <td className="level-right">
                        <span className="button is-uppercase is-small is-info">
                          {props.job.contact_name}
                        </span>
                      </td>
                    </tr>
                    <tr>
                      <td width="5%">
                        <i className="fa fa-bell-o"></i>
                      </td>
                      <td className="is-uppercase">Контактный email</td>
                      <td className="level-right">
                        <span className="button is-uppercase is-small is-info">
                          {props.job.email}
                        </span>
                      </td>
                    </tr>
                    <tr>
                      <td width="5%">
                        <i className="fa fa-bell-o"></i>
                      </td>
                      <td className="is-uppercase">Контактный телефон</td>
                      <td className="level-right">
                        <span className="button is-uppercase is-small is-info">
                          {props.job.phone}
                        </span>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}

export default JobBody;
