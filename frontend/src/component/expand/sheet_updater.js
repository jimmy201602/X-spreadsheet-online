//MODIFIED:上传表格数据到后端
import DataToJson from './data_to_json';
import {throttle} from './throttle';
import {Api} from './api';

export default function initSheetUpdater(toolbar){
  let sheetUpdaterDiv = document.createElement('div');
  sheetUpdaterDiv.className = "x-spreadsheet-toolbar-container-innerdiv";

  let sheetUpdaterBtn = document.createElement('button');
  sheetUpdaterBtn.innerHTML = "生成报表";
  sheetUpdaterBtn.className = "x-spreadsheet-toolbar-expand-btns-updater";

  sheetUpdaterBtn.addEventListener("click", throttle(function(){
    let data2josnInstance = new DataToJson();
    data2josnInstance.setSendDes(Api.createReportApi);
    data2josnInstance.send();
  }, 1000), false);

  toolbar.appendChild(sheetUpdaterDiv);
  sheetUpdaterDiv.appendChild(sheetUpdaterBtn);
}
