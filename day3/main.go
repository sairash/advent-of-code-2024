package main

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"
)

func main() {
	total := 0
	text := `
	'{}mul(339,896)>^+!)^mul(799,303)don't()>mul(188,763)'<};who()select()%;+mul(924,355)mul(492,757) what()mul(582,171)][*+select()#mul(840,899){!when()from()%<mul(711,51)when()why()} ~mul(131,623)&select()^how()mul(966,541)[*>where()mul(318,527)} :!-'mul(530,886)?}>mul(937,475) $;),%:}mul(201,723)where()select()mul(673,729)why()who()^'who()mul(673,694)[+mul(295,161)[!how(88,740)*mul(364,904)how()<]when()+where()mul(329,432)when()mul(499,11)who(238,444)<mul(533,879)'&who()#$;(&'<mul(65,49)#where(630,776)#mul(979,846)select()%]!<>)#~mul(775,866);,[)':where()%{[mul(835,890)+&&select()&[when()why(783,259) select()mul(735,871)!)when()'what()[/:mul(952,728)mul(633,505)@ -(?mul(176,469)*%what()>what()who()@{+do()'mul(117,634)-?(^^%:mul(234,514)where()@%mul(291,507)#from()*!*mul(668,282)@&)>,:select()>{%mul(195,300)-why()select()+&~>/^from()mul(801,834)why()</when()<&]mul(265,493)$what(382,576)#(+#']mul(590,771)%/mul(716,564)#}'mul(359,60)*~];#]mul(197,425)who()+^^?[:@[mul(752,102)]mul(271,88)mul(933,166)why()@,$^+?mul(343,220)+'what();mul(309,990){from(665,45)why(){ when(){ mul(782,953)+,:who()@]*mul(779,796)select()mul(616,478)&]>~mul(463,630){*, from()$}:@mul(280,83)when()[mul(358,910)[;'why()where()mul(242,569)from()#<>from()&mul(553,455)%who()<when()where()[mul(567,429)what()mul(257,307)}<don't()what()>)mul(284,63)%%*+?mul(437,226)* }how()when()~%'mul(57,491)]select(918,666)where()$when()why()'from()?]mul(321,301)'~:mul(619,356):mul(78,106)what()}!+~mul(609,442);  $where()$who()mul(996,918)mul(217,653)@##:#mul(998,408))~<#where()from()who()who()what()(mul(305,980)-~(:>where();when()#mul(721,412)how()'< { mul(143,735){:]why(){#),@mul(670,301)$when(),}why()]?why(839,544)mul(120,681){when()$[?@-)mul(805,510)>from()))when();?'#mul(104,633)%<$%}why()mul(555,387)@$+mul(850,237)!^where()<}from()select()from()<@mul(298,559)who():from()+what();mul(556,540)$%<&(%don't()$/':'*)(mul(976,624)!~*/%why()mul(790,645):~^from()[{+*!mul(153,86)+select(){#!from()how()$mul(980,956)>from()select()}<}@}?~mul(151,20)select()mul(703*(){+]who()what()mul(827,322)+](}mul(531,132what()where()+mul(933,2){&$how()%#;]don't())[]mul(845,519)how(),]when()^mul(518,563)#,++$#mul(500,591)(#/what()where()how()from()mul(243,908);mul(574,691)/who(),who()how()&mul#{where()when()]!@mul(534,43)}do(),}/from()when()~{&@mul(92%what()~}mul(496,669)^(!+ ^~mul(28,334)mul(621,688)]mul(627,561))mul(206,37)]~^&mul(288,740,<@mul(540,77)<&:who(594,229)&'*who(){mul(923,453)mul(733,228)where()how()mul(104,17)/!why()~what()*@}mul(500,830)#'(&%{select()*?mul(301,211)]>@@,mul(21,358) ?mul(285,542)how()from())mul(361,19)(who()%}select(){*mul(362,324)<[]'&when()'mul,why()mul(352,273)mul(742,91)>mul(624,723)) ;@+mul(14,149)(from()%%,(mul(547,492)~+mul(712?@@@&{{mul(972,531)
]&%where()~}who()[how()]mul(602,51)how()+&>,{>] #do()from()~{,*[-mul(862,742)how()why()]%mul(432,72)what(){:do()%@!}-mul(663/+,what()--(&?mul(384,302)'@(mul(649,348)+from()%mul(184,596)~+}~mul(719,53)mul(634,179){-:where()mul(684,320where()when(395,300){who()how()^/;where()mul(849,756)!mul(530,108)#*+}what()^(]select()mul(333,615)[why()%?]~$how()mul(314,366)}where()mul(222,364)<){*[mul(449,95){:who(844,554)<;why()$who() mul(831,201))$mul(408,650)who()what()}<[do()~how()select()!]'why()<when()mul(478,641)what()<mul>]how(289,983)* <+%&mul(836,460)%mul(339,868)why():from()from()%mul(91,296)!+^,*when()who()from()-$mul(6,37)when()when()mul(69,574)who(),from(),how()mul(431,678~+how()]mul(644,184)-(?why(571,97)])why()from()mul(516>select()mul(67,86)+~%^!~what();mul(526,440)!+>?<:&mul(81,534)&}'mul(64,25)[-;mul(828$<>*mul(157,667)@[ *who()mul(356,285)select()(~*do()how()';why()&^?mul(165,944)select()mul(980,979)<:!~%mul(15why()$ mul(109,665)&-!why()]<'&mul(887,673)]mul(906,700)#mul}@-where()/{{ -mul(935,960),)''[{mul(533,431)what())'@mul(63,509)@why(464,997)$]mul(164,971)select()~where()how()#>' when()>mul(301,62) +;'+what()}!->mul(722,492))!'mul(262,457):@when()-/mul(902,705)~#(mul(640,550)/*$$#select()where(905,349)!&&don't()when()mul(998,104)select()from()select()'when()mul(37,27)!where()$:do()}mul(160,45)mul(716,642~,{+&+!}[}mul(281,768)who()-?;);%mul(270,620) mul(793from()(![(! : who()mul(481,293)?mul(264,360)where(){from();(select()~!from()mul(748,940)[~]why()$[+how(709,453)mul(590&!+*why()]when()mul(182,631)(how()?(select():;&]{mul(83,366)%when()when()&mul(878,366)why()[:,]mul(77,997);%/$&%]mul(827,204)mul(919,654)>,where()%+mul(678,952)who()@select()}*(mul(344,894)where()mul(408,29)#*!{}*~where(906,182)mul(144,162)!&#select()how()&why()~*#don't()-()]~:how()mul(803,649)]@?#;mul(170,978)mul(263,974)!@why()$how()@mul(155,265)&/%^/mul(571,825)$where()mul(507,171)from()^(~*mul(437,680)from()who()>select()}mul(332,921) where()mul(218,74)})from()/mul(470,570)why()@?who()don't()@({*mul(931,767)mul(486,567):&])%/{]%mul(901,942),' ]why()^where()do(){#,what()mul(331,184)how()when()how()*{:^){mul(339,48){'(what(545,390)mul(818,891)who()mul(828,226), how()where()'#,%?mul(798,324))<how()mul(145,827)mul(256,218)who()?};when()/^mul(125,982)!%mul(274,98)}%what()-:who()},;mul(748,186)(()when()?where()why()<where()do()<%where()mul(556,171){>';mul(337,760)[#mul(350,889)~#how()mul(859,480){}^?&select()where()do()~<mul(808,237)$ % [#mul(218,295)mul(583,684);from():,from()+&-mul(628,340)[why()}why()?how()/,'mul(639,874)^}who()!why()mul(607,392)-&who()@+what()%mul(263,676)+%:]{ :-select()mul(166,500)/mul(711,477) : {how()<mul(939,832)'+<{!,when(),?)mul}[what()where()+:-,from()mul(457,751)&from()+why()mul(219,492)^/$! $from(832,913)@+!mul(892,437)+[>); who()^mul(640,455}why()#mul(744,51)'[who()/> select()mul^where()select()^}^mul(450,596))select()what()}&%;?mul(218,957)+*who(),}do()from()^} when()+[select(938,490))select()mul(406{how()(+who()()select()mul(329,937)&!mul(693,766)<{+}<[@mul)]%why()when()){~[ who()mul(888,144)~$:,mul(517,97)@) ~mul(394,320)why()when()who()(,%mul(761,855)
mul(22,362)('from(886,421)]mul(730,655)[@,how()(mul(692,165)]&$when()!}from()%mul(481,375)where()~mul(954,570)?why()-+mul(338,656)who()~ <}from()mul(616,31)where()/]:select())?from()mul(113,2):?$mul(295,905)];mul(410,181)@%${^how()>select()where()-mul(779>+what();^who(),>)mul(599,200)%~][select()>+>mul(486,481)*!who(693,495)-$mul(237,686)? how()! -@#do()#<(where()-&'>&what()mul(321,434)}@what()~/from()do()?who()$where()mul(328,792)select()how()mul(82,296)#</ why()what()}mul(859,169)[>,who())/when(637,168)mul(465,709)mul(208,775)^[@when()>>##<>mul(379,29)%mul(826,43){when()?who()*why()do()&[):@mul(411,966)^mul(24,557)<;where()mul(391,794)#;mul(592,819)+,}'%'mul(210,928)%mul(29,613);$who()why()!]who()mul(56,646)*@]-{~+:mul(425,457)>mul(896,578)%(how()](*where()when()select(237,23)mul(895,482)~<{mul(432,547)who(471,124){mul(483,785)*mul(422,876)^>& ;(^where()~#mul(709,114)(:;where()select()?%mul(263,276)&?;from()&&~(mul(113,694)who()mul(228,70): >[:@ ;@mul(707,104);:mul(423,229)&$[]>who()mul(194,895)><&&when()%%mul(836,144)<^!~)/;#who()>mul(786,723)?!#[mul(287;why()mul(734,761)!who()/<)mul(520,746)where(){!>>select()how()mul(185,986)mul(566,786)why()when()[~do(),}mul(188,610)/+^%<why()mul(784,533);-;?#/when(743,88)mul(841,352)from()why()how()+/!<mul(728,38);<why()from()>-how()[mul(671,105)*[}[mul(403,996)mul(214+($?{when()mul(268,651),[>mul(660,864)%-/how()->([~mul(769,53)?from(197,675)^[-mul(83,519)where()select()don't()[:from()@{who() ?mul(305,335)[when()when();where(751,621)what()mul(395,86)how()?,who():>mul(349,362)how()?*select()when()from()who()where()''mul(414,725)*)when()select()]+mul(180,197)$who()why()&%'}mul(531who()#{mul(370,295)who()%mul(121,586)*^^%?>{,when()mul(944,189)&[/)select()>&^mul(222,28)@</<mul(752,302)}!<why()'$where()?,mul(974,242)}>!where()'mul(449,827)^[mul(289,78)$how(287,947)mul(337,811)why()''-what()[/when(370,472)from()>mul(865,636){#mul(524,198)why(714,875)!*%mul(181,23)^)why()?:what(630,704)+}mul(569,165);when()')where(597,70):$where()why(),mul(15,411)* ) &don't()mul(124,709)$;/[+select()(</mul(99,652)mul(53,14),select()where()mul(380,904)^}!?**[@when()mul(491,229)$#where()mul(245,344)select()select(){mul(297,527)'<from()/-,>{mul(50,277))*;+*<#do()!@mul~mul(19,630),,,*+{mul(404,379)mul(72,663)when()don't()where(221,302)from()^>mul(942,445)+^from()]from()*{)mul(83,601)-+what(): where()what()from(),mulwhy()where(856,731)&/+mul(777,574)!+when()<where()why()'>don't()*^<what()from();>mul(680,66)how()$mul(361,449):,how(766,248)#}&}[mul(869,603);;where()what()mul(385,816)[!' <-~from();mul(298,605*<mul(189,109)from()+/)!:mul(451,205)mul(949,138)from()+# $;}mul(356,99)who()*#-)':mul(95,448)who()$@who()-%&mul(167,343)mul(300*~%who()$(how();}])mul(919,379)>from()^mul(573,375)when()@(where()%mul(871,907)mul(718,918)?mul/~ ~what()-!select()~do()why():mul(682why()]mul(585,886)(,?+*?%!mul(684,834)what(786,470)mul(443,590))where(228,285) !/%?mul(815,879)#!/usr/bin/perl@mul(444,941)$select(687,764)%'(where()>-/mul(180,328)
[/:@how()*what()!+mul(911,368)?/what()~(+]mul(843?$,who()mul(865,234)@]/-from()mul(397,906)^mul(806,349)]how()where()^)+%select()when()~mul(827,131)]don't()@+/mul(44,818)[,<,mul(295,441)what()/select()]^mul(756,90) [mul(67,416){mul(230,994)select()how()who()/mul(66,226);<~!when()mul(325,467),mul(6,370)mul(619,21)~<what(761,805)who())+?;^-mul(5,165{don't()how()who()#,mul(713,804) mul(737,356);mul(905,649who()what()mul(139,324)~from()mul(502,936)when()-&select()^']]who()mul(404}>mul(699,668))^<from()%!%>:mul(755,644)<%?#)mul(46,923!mul(730,880)]~]how()-why()'&mul(952,543)from()<what()what()where()>>!-mul(123,880){(<when()<mul(425,371)from(),how()+?]-*!where()mul(425,817)?!mul(668,3) who()mul(785,430)*mul(607,686)?] $~mul(979,796)]/@'why()mul(244,801)+when()how():mul(311,17)@:,#-why()($mul(486,732)mul(480,165)mul(153who():when()&~)who()$mul(662,582)<((what()* {/(&mul(894,455);;(how()what()[>mul(555,437)*mul(692,73)&*when()~>mul(465,602)(mul(471,204){%from();don't()$mul(945,735)from()select(520,626)>@who()]who()mul(615,73)):(;^^when(793,925)*&do(){mul(431,683)*+#select()where()from()?+mul(254,617)#%where()>;%don't()$who()#how()^[how()why()*mul(907where()select()(!:'!?@mul(208,995)}/:when()how(415,229)^-'from()mul;mul(22,79),']^?(mul(583,536):mul(355/-where()?<mul(281,314)how()^!+<;>~%,mul(990,358)//+&how() do()?who()%!}]mul(603,599)@mul(285,652){&@@mul(808,857when()when()%select()/'$@mul(585,541)from()@mul(136*#from()@mul(710,522) #*when()*}/mul(801,485) >/mul(393,477)where()(mul(13,599)what()when()(*%>@?];mul(808,562)>mul(407,85)mul(244/$&!+where()mul(67,663)<from()-{where()% where()mul(629,684){}#^)-why()where()+mul(79,607)*don't()'from(27,368)*where()<#^mul(697,649)/(why()~from()*mul(448,917)mul>-mul(934,570)]mul(857,473)who()mul(585,495)where()mul(45,904)!when()where())-:mul(747,283)why()#where();what()how(){'from()mul(405,574)[,?what()why()-([mul}where()select(450,140) /<mul(198,934)when(394,203)why()<']@,mul(299,635)>@who():<mul(629,260)&!/!who()mul(360,191)select()#*'mul(409,799)select()}} *from()*,mul(902,917)do()where()why()%}}]+mul(548,522)^]how()->@who(),>~mul(104,734)$]#-#who()mul(760,886)what()<where()[#mul(150,972)mul(276,427)from(729,212)where()% {!mul(534,660)when()-?mul(406,3),!%*why())mul(990,129)how() what()mul(532,895)how()when()mul(869,39))/;->&where(352,510)->(don't()mul(863,264)$<)where()*/(@'mul(756,795)^)]mul(278,155):&!%$select()mul(189,750)[$#-/mul(549,580))^how(152,70):$mul(28,530)-],;]mul(33,157);!/+?what(253,786)%what()mul(841,40)+&when()why()^mul(898,936)!])mul(891,523);>mul(312,16)@how()*where()'where()<@?mul(967,420)/}why()~,mul(581,636)/ [mul(673,139)who()>mul(578,980):,<what():}{&: mul(605}^-+~]'@$ }mul(229,41)#@)+mul:mul(447,836)how()*%*what():;['*mul(672,963);what()mul(287,244)%;+<%{@mul(448,425)//+~([who()mul(871,92)]}?++[--%*mul(42,503)#?[{$^}>%mul(75,107)+)how()-mul:/where()why()mul(315,687)!{%'what()mul(110,111)+ #:%!mul(731,760)
(+&mul(887,468)$::)],mul(765,973)'from()from()*mul(810,344)?what()mul(768,468)~'+)select()where()(select()where()$mul(576,358)%??'mul(41,789how()when()&-mul(606,191)!when()'~mul]]how():;~{how()mul(15,34)>%%*^how(54,122)$@mul(739,223)how()~*@!don't()]:{~'@>why()<mul(929,311)(%@;mul(949,785)>-]&how() **;}mul(900,428)select()]$'mul(874,363)what()@mul(892,45)^*+mul(387,178)?<how()mul(383,479)/@#who()/)do()^}where(662,769)~$mul(845,224)$from()mul(60,887)mul(773,136)%^mul(436,490)%+mul(283,346);?mul(77,681)^##mul(556,520)$how()(&^how()select()) who()mul(418,701)!^:},where()/%mul(871,886)when()[mul(409,599)*from(463,269){from(324,295)from()who()mul(790,739)mul(677;why()}^when()how()+from()don't()&*%,why()#!how(252,857)-?mul(728,703)how()who() ')^from()'how()^do()[mul+!@;~select()select()mul(108,547)select()mul(96,190)];,{what()/when(742,998){})mul(632,699)how()'(what()when(),^mul(917,127)where()^()<mul(763,236)/^@+(when()^!%*mul(270,488))?what(116,413)!?:mul(922,831)$,who()@!*(>mul(823,845){where()mul(854,982);;/how()-why()mul(899,363)[where()who()what()>$]['do()<'how())mul(201,507)select(): >select()select()what()/where()*do()}@when(392,773)?mul(231,610){- why()*from()when():select():mul(334,751)from()how()~+who()-mul(811,647){^mul(116,805)^where()@}mul(691,631) when()&%from(),@^mul(336,461) ,what())who()when(540,382)'mul(549,430)%]from())@mul(339,808)?mul(264,497)'when()what()~who()~@how()$+mul(965,916)('who())%~from()from()%mul(776,506)/select()mul(385,184)*##select()mul(691,451)]$$mul(303,437);!when()&/<>]^mul(524,315)#^)mul(42,992who()}&-select()mul(902,182)(!''where()'+mul(48,755)@~what(644,7)$select()&%who()-mul(629,650)!^mul(822,985){/select()why()where()!who()%%mul(102,630)why()-{don't(),mul(166,527)'where()mul(245,921)select()-select()mul~>select()'%who()<+,mul(795,941))%$,what()where()>mul(414,585)from(26,999){mul(293,208)when()?#?<!who()mul(781,159)/#do()select()mul(409,682)/select():who()mul(565,948)mul(903,713):how()>mul(99,672)why()(select()+]when()mul(300,836)[<mul(129,851)mul(545,309)why(388,433)@~/mul@how())&*mul(466,194) };;@what()who()(mul(213,565what()why()how()#%,<where()mul(844,768))&()]}-<why()+don't()/};'</&mul(134,385)mul(176,788)[-]]mul(111,769)[;{mul(796,855)when()?$who()?,;/$mul(800,759)&what()+why()<how()*mul(20,432)why(611,106)where()* /]?$%do()who()/!^mul(691,133)who();^(:>}where()/mul(425,328))mul(702,532)?#(mul(45,856) ]/$mul(220,616)^when(){^* ?where():,mul(931,398)$;-mul(471,783)why(){:who()what()]>mul(276,590)select()when() /$(^<'mul+*&do()&)>;<mul(675,852)~;mul(128,569)!who()from(306,117){-what()mul(868,808)('why()mul(71,753)(why()~mul(250,975)?how(){;select()->mul(785,175)mul(131,573),]from():(mul(833,970)! ><<(mul(496,285)&from()~select(){mul(296,374)?[#&from()from(){{when() mul(718,993)@who()&>mul(639,708)why()where()/-<{[ how()&mul(187,633)#:]:?mul(872,562)/who()[-who()(>$%}don't()!{:*mul(82,739)select()]+ ?$?@when()mul(830,429)<what() !#^where(985,664)don't()$:,;}}where()mul(920,165)/,&:#:+'?^mul(979,57)what(740,146)},<;:[+from()-mul(481,113~when()( ~why()[,how()don't()where()^from()%from()mul(845,384)]why(),mul(884,267%what()how()why()where(358,396)mul(795,514):<what()select()'when()-mul(415,325){select()]where()&mul(876,97)@}#@^[mul(204,271)
$'do()mul(983,642);>+;:$>mul(390,30)(&@%&'} {%mul(208,444)mul(854,207);:@where(774,785)mul(120,222)mul(885,372)$ '<[mul(476,77)select()from()mul(305,758)}#[>;,@where()~<mul(999,999)!&?mul,why()when()how()&from()+why(610,462)mul(28,578)+-,&[>how()(why(){mul(2,348)%@(~}how():mul(148,153)where()^:'why(907,374)+]mul(375,986);{]where()(!@what()]:mul(254,345))*)where()select()@?^#what()mul(94?what()!@when()select(),$?select()mul(966,420)select()/+{[;!(>what()mul(688,942)&]'^)'#!mul(363,573):$,from()select()mul(260,171)#mul(116,728)'(?mul(51,309)[!^mul(400,128))select()$-^how(740,875)where()@from()mul(319,269)select(844,13)who()#/mul(431,542)mul(794,709)- who()*do()*:)select()mul(12,579)@%])[what()&/mul(361,146)why()>:'-^&;mul(465,576))select()@%mul(101,476)who()%}' [#mul(13,38);:how()from()from()#}#&-mul(30,350)what()/+select()+]don't()!#:@;,[when()mul(303,869)when()where()]/'^!^ mul(938,614) {who()?!what(363,886)mul(439,873)>who()what()-+when()where()[why()mul?from()what()> where()who()!mul(510,226) {how()mul(353,498)]select()why()?[why()'*}(mul(709,649)~?select()*($>[from()mul(144,790)!%!mul(653,286)[<$)*+from()what(),from()mul(373,21))] select()?;+mul(601,965)what(), ,:mul(970,654)~why()(~*<+&mul(19,700)what()$what() <{<mul(850,749)when()/select()#;*&mul(561,720)from()mul(439,509)who()from()@}/,mul(390,146)who(888,573);&;@+:mul(658,554)don't()who()*how()}where(){+how()&mul(769,268)^from()?;when(518,375):'mul(570,233) )[?from()$ *mul(183,547)what(277,464) !;;mul(490,847);+)#?*where()why()how()mul(23,467)(:!mul(305,784)mul(237,433)+~[who()<:-$mul(842,162)?>({,)'how()-@mul(629,950)how()when()who()what(623,556)mul(891,30)(!don't()>who()who()%+ [)when()mul(808,962)+?</}*,&^@mul(90,782)[[mul(944,224)mul(442,115)from()where();when()+mul(162,476)mul(689,801)%><({~$#+do()}what()how()(what()*-@mul(296,108)>]^+&mul(592,463)~['%%& mul(733,447),#how()where()select()how()who()>mul(496,360)(&%*{+what()!mul(615,52)#<why(665,627)?;]$:[mul(865,200)') ')}@{}@mul(275,101)where():@</<#, select()mul(693,721)/who();from()>&*,mul(807,92)&]mul(544,513){/+who(868,321)from()mul(164,401)how()][%what()where()(+[>mul(815,703)from()(when(),*select() {<mul(20,330)mul(78,579)from(){{where()]+-*#&mul(141,703)mul(832,264)$/do()+/##>what()mul(437,310)who(226,447)mul(45,389)#&who()<mul(672,921)what()mul(9,140);)mul(19,667)};]+@why()[,)(mul(256,632)#!'!$;$where()how()mul(648,453)+'who()(?!}mul(453,188)!from()$ ]when()what(326,8)mul(96,408)mul(291,868)%~where()when()]#&^%mul(899,304)'>*[mul(659,114))what()#:+[%mul(686,605)<',,{who()^@-mul(127,293)& ;@mul(608,869)?%+*select(74,799)who()when()/^select()don't()where()$why()>why()when()%mul(752,203)-}'{^#;[%^mul(307,633)%when(970,30)mul(265,251)what(42,790)#mul(188,777)<:*& when():/when()?mul(94,809)mul(621,327)+;/*,,[select();~do()^(when()<what(554,336)mul(178,771)+%<{;;when(),/>mul(896,407),$mul(280,745)*:){^~:({where(796,413)mul(262,847)why()&&<where()$~}#mul(792,379):$!-!select()/]mul(199,174)-}!%#]mul(639,740)'select()@from()&[#

`
	text_input_chan := make(chan (string))
	got_value := make(chan ([]string))

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			v_t, ok := <-text_input_chan

			matches := re.FindAllStringSubmatch(v_t, -1)

			for _, match := range matches {
				got_value <- match
			}
			if !ok {
				close(got_value)
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			v_t, ok := <-got_value

			if !ok {
				break
			} else {
				num1, err1 := strconv.Atoi(v_t[1])
				num2, err2 := strconv.Atoi(v_t[2])

				if err1 == nil && err2 == nil {
					total += num1 * num2
				}
			}
		}
	}()

	go func() {
		enabled := true
		previous_found_text := ""
		for x := 0; x < len(text); x++ {
			previous_found_text += string(text[x])

		}
		if enabled {
			text_input_chan <- previous_found_text
		}
		close(text_input_chan)
	}()

	wg.Wait()

	fmt.Println(total)
}
