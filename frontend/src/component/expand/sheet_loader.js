//MODIFIED:加载json数据，然后显示在表格中
import JsonToData from './json_to_data';
import {throttle} from './throttle';
import {Api} from './api';

export default function initSheetLoader(toolbar){
  let sheetLoadDiv = document.createElement('div');
  sheetLoadDiv.className = "x-spreadsheet-toolbar-container-innerdiv";

  let sheetLoaderBtn = document.createElement('button');
  sheetLoaderBtn.innerHTML = "加载";
  sheetLoaderBtn.className = "x-spreadsheet-toolbar-expand-btns-loader";

  sheetLoaderBtn.addEventListener("click", throttle(function(){
    let json2dataInstance = new JsonToData(Api.getReportDataApi); //这里创建实例会从后端接口读取数据到实例属性中，然后在异步访问后自动执行加载
    json2dataInstance.gather();
  }, 1000), false);

  toolbar.appendChild(sheetLoadDiv);
  sheetLoadDiv.appendChild(sheetLoaderBtn);
}
