class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            companies: []
        }
        // this.serverRequest = this.serverRequest.bind(this);
        console.log(this);
        console.log(this.state.companies);
    }

    componentDidMount() {
        fetch("http://localhost:3000/api/companies")
            .then(res => res.json())
            .then(companies => this.setState({companies}))
    }    
    
    // serverRequest() {
    //     $.get("http://localhost:3000/api/companies", res => {
    //         this.setState({
    //             companies: res
    //         });
    //     });
    // }

    render() {
        const companies = this.state.companies;
        return (
            <div className="container">
                <div className="col-lg-12">
                    <br/>
                    <h2>GovtechList</h2>
                    <div className="row row-eq-height">
                        <div className="panel-body">
                            {companies.map(company =>
                                
                                    <div className="panel panel-default">
                                        <div className="panel-heading"><a href={company.official_url}>{company.company}</a>（{company.founded_year}）</div>
                                        <div className="panel-body">{company.description}</div>
                                    </div>
                                
                            )}
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

// class Company extends React.Component {
//     constructor(props){
//         super(props);
//         this.state = {
//             liked: "",
//             companies: []
//         }
//         // // this.like = this.like.bind(this);
//         // this.serverRequest = this.serverRequest.bind(this);
//     }

//     // like(){
//     //     let company = this.props.company;
//     //     this.serverRequest(company);
//     // // }
    
//     // serverRequest(company) {
//     //     $.post(
//     //         "http://localhost:3000/api/companies/like/" + company.id,
//     //         { like: 1 },
//     //         res => {
//     //             console.log("res... ", res);
//     //             this.setState({ liked: "Liked!", jokes: res} );
//     //             this.props.companies = res;
//     //         }
//     //     );
//     // }

//     render(){
//         return(
//             <div className="col-xs-4">
//                 <div className="panel panel-default">
//                     {/* <div className="panel-heading">
//                         #{this.props.company.id}{" "}
//                         <span className="pull-right">{this.state.liked}</span>
//                     </div> */}
//                     <div className="panel-body">
//                         {this.props.company.company}
//                     </div>
//                     {/* <div className="panel-footer">
//                         {this.props.company.likes} Likes &nbsp;
//                         <a onClick={this.like} className="btn btn-default">
//                             <span className="glyphicon glyphicon-thumbs-up"></span>
//                         </a>
//                     </div> */}
//                 </div>
//             </div>
//         );
//     }
// }

ReactDOM.render(<App />, document.getElementById('app'));