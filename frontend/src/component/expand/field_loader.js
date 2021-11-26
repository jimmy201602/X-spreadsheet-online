//MODIFIED:加载数据源的字段
import FieldListener from './field_listener';
import {throttle} from './throttle';
import {Api} from './api';


export default function filedLoader(toolbar){
  let fieldLoaderDiv = document.createElement('div');
  fieldLoaderDiv.className = "x-spreadsheet-toolbar-container-innerdiv";

  let fieldLoaderBtn = document.createElement('button');
  fieldLoaderBtn.innerHTML = "加载字段";
  fieldLoaderBtn.className = "x-spreadsheet-toolbar-expand-btns-fieldloader";

  fieldLoaderBtn.addEventListener("click", throttle(function(){
    let fieldListenerInstance = new FieldListener(Api.loadTableColumnApi);

    // 目前是针对postgreSQL查询写死的数据
    let postBody = window.xs.currentField.source;

    fieldListenerInstance.setBody(postBody);
    fieldListenerInstance.send();
  }, 1000), false);

  toolbar.appendChild(fieldLoaderDiv);
  fieldLoaderDiv.appendChild(fieldLoaderBtn);
}

