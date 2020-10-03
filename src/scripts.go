package main

var FourDDownload = `SET('PG', WC_MethodPage('http://www.singaporepools.com.sg/DataFileArchive/Lottery/Output/fourd_result_draw_list_en.html?v=123', 'GET', '', ''));
SET('URLS', SO_TagMatch(GET('PG'), 'queryString=''', ''''));
FOR('VAL', GET('URLS'),
   SET('FL', GET('VAL'));
   SET('U', JOIN('', 'http://www.singaporepools.com.sg/en/4d/Pages/results.aspx?', GET('FL')));
   SET('PG2', WC_MethodPage(GET('U'), 'GET', '', ''));
   SET('NO', SO_TagMatch(GET('PG2'), 'Draw No. ', '<'));
   IF(SO_FileExist(JOIN('', '{PARAM0}', '\\' '4D-result-', GET('NO'), '.txt')), '=', '1', EXIT(), PASS());
   SET('Prizes', SO_TagMatch(GET('PG2'), '<td class=''tdFirstPrize''>', '</td>'));
   SET('Prizes', GET('Prizes');SO_TagMatch(GET('PG2'), '<td class=''tdSecondPrize''>', '</td>'));
   SET('Prizes', GET('Prizes');SO_TagMatch(GET('PG2'), '<td class=''tdThirdPrize''>', '</td>'));
   SET('Starter', SO_TagMatch(GET('PG2'), '<tbody class=''tbodyStarterPrizes''>', '</tbody>'));
   SET('Prizes', GET('Prizes');SO_TagMatch(GET('Starter'), '<td>', '</td>'));
   SET('Console', SO_TagMatch(GET('PG2'), '<tbody class=''tbodyConsolationPrizes''>', '</tbody>'));
   SET('Prizes', GET('Prizes');SO_TagMatch(GET('Console'), '<td>', '</td>'));
   FOR('Prize', GET('Prizes'),
      SO_AppendToFile(JOIN('', '{PARAM0}', '\\', '4D-result-', GET('NO'), '.txt'), GET('Prize'));
   );
);`
