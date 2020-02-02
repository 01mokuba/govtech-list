import React, { Component } from "react";

class Company extends Component {
  render() {
    const { company } = this.props.company;
    console.log(company);
    return (
      <div className="panel panel-default">
        <div className="panel-heading">
          <a href={company.official_url}>{company.company}</a>（
          {company.founded_year}）
        </div>
        <div className="panel-body">{company.description}</div>
      </div>
    );
  }
}

export default Company;
