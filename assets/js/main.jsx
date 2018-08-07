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


class News extends React.Component {
  render() {
      const { data } = this.props
      let newsTemplate
      if (data.length > 0){
          newsTemplate = data.map(function(item){
            return (
              <div key={item.id}>
                  <p className="news__author">{item.author}:</p>
                  <p className="news__text">{item.text}</p>
              </div> 
          )
          })
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