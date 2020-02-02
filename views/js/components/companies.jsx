import React, { Component } from "react";
import Company from "./company";

class Companies extends Component {
  render() {
    const companies = this.state.companies;
    return (
      <div className="container">
        <div className="col-lg-12">
          <div className="row row-eq-height">
            <div className="panel-body">
              {this.props.companies.map(company => {
                <Company key={company.id} company={company} />;
              })}
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default Companies;
