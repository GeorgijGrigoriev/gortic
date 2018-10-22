const MyNews = [
  {   
      id: 1,
      author: "Test",
      text: "Test text"
  },
  {   
      id: 2,
      author: "Tester",
      text: "Test text acepted"
  }
];


class Article extends React.Component {
  render() {
      const { data } = this.props
      
      } else {
          newsTemplate = <p>К сожалению, новостей нет. </p>
      }
      return(
          <div className="news">
              {newsTemplate}
              {
                  data.length ? <strong>Всего новостей: {data.length}</strong> : null
              }
          </div>
      )
            
}
}

const App = () => {
      return (
          <React.Fragment>
              <News data={MyNews} />
          </React.Fragment>
  )
}
ReactDOM.render(
  <App />,
  document.getElementById('root')
);