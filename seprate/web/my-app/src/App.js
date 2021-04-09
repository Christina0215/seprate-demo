import './App.css';
import { useState } from 'react';
import axios from 'axios';

function App() {
  const [num,setNum] = useState({
    a:0,
    b:0,
  });
  const [ans,setAns] = useState(0)
  const cal = ()=>{
    axios.get('http://localhost:8000/?'+'a='+num.a.toString()+'&b='+num.b.toString()).then((response) => {
      console.log(response)
      setAns(response.data)
    });
  }
  return (
    <form>
      <input name="a" value={num.a} onChange={e=>{setNum({...num,a: e.target.value})}}></input>
      +
      <input name="b" value={num.b} onChange={e=>{setNum({...num,b: e.target.value})}}></input>
      <button type="button" onClick={cal}>提交</button>
      ans={ans}
    </form>
  );
}

export default App;
