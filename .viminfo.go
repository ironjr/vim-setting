# ì´ viminfo íŒŒì¼ì€ ë¹”ì´ ë§Œë“  ê²ƒì…ë‹ˆë‹¤ Vim 7.4.
# ì¡°ì‹¬ë§Œ í•œë‹¤ë©´ ê³ ì¹  ìˆ˜ë„ ìˆìŠµë‹ˆë‹¤!

# ì´ íŒŒì¼ì´ ì €ì¥ë˜ì—ˆì„ ë•Œì˜ 'encoding'ì˜ ê°’
*encoding=utf-8


# hlsearch on (H) or off (h):
~h
# Last Search Pattern:
~MSle0~/core

# Last Substitute Search Pattern:
~MSle0&block

# ë§ˆì§€ë§‰ìœ¼ë¡œ ë°”ê¾¼ ë¬¸ìì—´:
$

# ëª…ë ¹ ì¤„ íˆìŠ¤í† ë¦¬ (ìƒˆê²ƒë¶€í„° ì˜¤ë˜ëœ ê²ƒ ìˆœ):
:wqall
:w
:wall
:wq
:q
:qall
:noh
:o src/timeframe.go
:o src/context.go
:'<,'>s/block/bx/g
:'<,'>s/blockchain/bc/g
:w\
:%s/log.Printf/t.Logf/g
:s/log.Printf/t.Logf/g
:s/\.log.Printf/t.Logf/g
:'<,'>s/\.model//g
:'<,'>s/Int/Float/g
:'<,'>s/)/, 10, 64)/g
:'<,'>s/"/Str/g
:'<,'>s/args/strconv.ParseInt(

# ì°¾ì„ ë¬¸ìì—´ íˆìŠ¤í† ë¦¬ (ìƒˆê²ƒë¶€í„° ì˜¤ë˜ëœ ê²ƒ ìˆœ):
? \s\+$
?/core
?/Topology
? block
? blockchain
?/network
?/Hello
? log.Printf
? \.log.Printf
?/log\.
? \.model
?/\.model
?/.model
?/ev
? Int
? )
? "
? args
? ok
? }

# í‘œí˜„ íˆìŠ¤í† ë¦¬ (ìƒˆê²ƒë¶€í„° ì˜¤ë˜ëœ ê²ƒ ìˆœ):

# ì…ë ¥ ì¤„ íˆìŠ¤í† ë¦¬ (ìƒˆê²ƒë¶€í„° ì˜¤ë˜ëœ ê²ƒ ìˆœ):

# ì…ë ¥ ì¤„ íˆìŠ¤í† ë¦¬ (ìƒˆê²ƒë¶€í„° ì˜¤ë˜ëœ ê²ƒ ìˆœ):

# ë ˆì§€ìŠ¤í„°:
"0	LINE	0
	/**
	 * delay_model.go
	 *
	 * Jaerin Lee
	 *     Seoul National University
	 */
""1	LINE	0
	
"2	LINE	0
	
"3	LINE	0
	
"4	LINE	0
	<<<<<<< HEAD
"5	LINE	0
	=======
"6	LINE	0
	>>>>>>> core
"7	LINE	0
		ProofOfWorkType ConsensusType = 0
"8	LINE	0
		"core"
"9	LINE	0
	
	V
	
"w	CHAR	0
	3238
<wjkhpjjbbwbhp:w€kd€kd€ku€ku€ku€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€ku€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kd€kd€kd€kd€kr€kri"MinDelay:w€ku€ku€ku€ku€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kl€kl€kldwimatp[stirring]:w€kd€kd€kd€kd€kl€kd€kl€kl€kl€kl€kl€kl€ku:q€kd€ku€ku€kl€kd€kl€kr€kri"mMeanDelay":w€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kl€kl€kl€kl€kl€kl€ku"Minuu:q€kl€kri"Minn€kd€kl€kl€kl"Max€kd€kd€kd€kd€kd€kd€kr€kr€kr€kr€kr€kr€kr"MinMultiplier€kd€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl"Mean<ul<ulMultipleierier:w€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kl€kd€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€kd€kd€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€krdwimap[string]€kd€kd€kd€kd€kd€kl€ku€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€krDelay"Delay€kd€kd€kd€kd€kd€kd€kl€kl€kl€kl€kl€kl€kl"Min€kd€kl€kl€kl"Max":w€kl€kd€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€krwwwwwwwwdwimap[stirngrintring:w€kl€kl€kl€kl€kl€ku€kr€kr€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kd€kl€kd€kd€ku€kl€ku€kd€ku€ku€kd€kl€kr€kr€kr€kr€kr€kr€krd€krim€kr€kd€kl€kl€klm€kd€kd€kd€kr€kr,mm€kd€kdm€ku€kr€kr€ku€ku€ku€kr€ku€kr€kd€krmodel€ku,pmodel€ku€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kd€kumodel:w€kl€kd€kl€ku€ku€kd€kl€ku€ku€kd€kl€ku€kl€ku€ku€ku€ku€ku€kr€ku€krwwwwwww€kr€kr€kr€kr€krdwimap[string]€kd€kd€kd€kl€kd€kl€kl€kl€kr€kr["sizSziozeze":w€kl€ku€ku€ku€kd€kd€kd€kd€ku€kl€ku€kl€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€kl€ku€kl€kd€kr€kr€kd€kr€kuwwwwwwwwwww€kr€kr€kr€kr€krdwimmap[string]€kd€kd€kl€kd€ku:w€kl€ku€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€ku€ku€kr€kr€krd€krim€kd€kl€klm€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€krmodel€kd€kl€kl€kl€klmodel€ku€kd€ku€ku€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€krmodel€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€krmap[string€kd€kd€kd€kd€kd€kl€kl€kl€kl€kl€kl€kl€kl€ku€kl€kl€kl:"""MainxinxX"€kd€kl€kl€kl€kl€kl"MaxX"€kd€kl€kl€kl€kl€kl"MinY"€kd€kl€kl€kl€kl€kl"MaxY"€kd€kd€kl€kl€kd€kd€kl€kd€kd€kl€kd€kl"MaxR€kr€kd€kd€kd€ku€ku€kd€kd€kd€kl€kd€kl€kl€ku€kl,m€kd€kd€ku€ku€ku€ku€ku€ku€ku€ku€kum€ku€ku€ku€ku€ku€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€ku€kl€klm€kd€kr€kr€kr€kr€krm€kd€kd€kd€kl€kl€kl€kl€klm€kd€kd€kr€kr€krm€ku€kr€klm€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€klm€kr€kr€kr€kr€kr€kr€krm€kdm€kl€kl€kl€kl€kl€kl€klm€kd€kd€kd€kd€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€kl€klm€kd€kr€kr€kr€kr€kr€krm€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku€ku:w€kl€kd€ku€ku€kl€kd€kd€kd€kd€kd€kd€kd€kd€kd€ku€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kd€kr€kr€kr€kr€kr€kr€kr€kr€kr€krim€kd€kr€kr€kr€kr€krm€kd€kd€kd€kd€kl€kr€krm€kd€ku€kl€ku€kl€ku€kl€kd€kr€kl€kl€kl€kl€klm€kd€kd€kd€kdm€kd€kd€kl€klm€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€kr€krm::w
"-	CHAR	0
	:

# íŒŒì¼ ë§ˆí¬:
'0  20  18  ~/.go/src/blockchain-simulator/src/txgen/model.go
'1  26  0  ~/.go/src/blockchain-simulator/src/txgen/model.go
'2  8  0  ~/.go/src/blockchain-simulator/src/core/simulator.go
'3  14  0  ~/.go/src/blockchain-simulator/src/core/timeframe.go
'4  167  0  ~/.go/src/blockchain-simulator/NERD_tree_1
'5  183  0  ~/.go/src/blockchain-simulator/NERD_tree_1
'6  70  0  ~/.go/src/blockchain-simulator/src/blockchain/node.go
'7  12  0  ~/.go/src/blockchain-simulator/src/core/iterator.go
'8  225  0  ~/.go/src/blockchain-simulator/NERD_tree_1
'9  50  16  ~/.go/src/blockchain-simulator/src/core/blockchain_test.go

# ì í”„ëª©ë¡ (ìƒˆê²ƒì´ ë¨¼ì €):
-'  20  18  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  1  12  ~/.go/src/blockchain-simulator/src/txgen/txgen.go
-'  35  0  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  26  0  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  43  0  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  48  0  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  46  0  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  51  0  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  32  0  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  31  15  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  2  7  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  1  0  ~/.go/src/blockchain-simulator/src/txgen/model.go
-'  38  0  ~/.go/src/blockchain-simulator/src/network/delay_model.go
-'  94  2  ~/.go/src/blockchain-simulator/src/network/network.go
-'  10  0  ~/.go/src/blockchain-simulator/src/network/delay_model.go
-'  1  0  ~/.go/src/blockchain-simulator/src/network/delay_model.go
-'  8  0  ~/.go/src/blockchain-simulator/src/core/simulator.go
-'  6  0  ~/.go/src/blockchain-simulator/src/network/coord.go
-'  7  0  ~/.go/src/blockchain-simulator/src/blockchain/event.go
-'  31  0  ~/.go/src/blockchain-simulator/src/core/simulator.go
-'  14  0  ~/.go/src/blockchain-simulator/src/core/timeframe.go
-'  1  11  ~/.go/src/blockchain-simulator/src/core/iterator.go
-'  1  11  ~/.go/src/blockchain-simulator/src/core/context.go
-'  12  2  ~/.go/src/blockchain-simulator/src/network/network.go
-'  57  0  ~/.go/src/blockchain-simulator/src/network/network.go
-'  11  0  ~/.go/src/blockchain-simulator/src/core/context.go
-'  13  0  ~/.go/src/blockchain-simulator/src/core/context.go
-'  15  0  ~/.go/src/blockchain-simulator/src/core/context.go
-'  6  33  ~/.go/src/blockchain-simulator/src/consensus/consensus.go
-'  7  0  ~/.go/src/blockchain-simulator/src/consensus/consensus.go
-'  19  5  ~/.go/src/blockchain-simulator/src/core/iterator.go
-'  167  0  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  181  2  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  1  0  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  183  0  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  194  2  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  70  0  ~/.go/src/blockchain-simulator/src/blockchain/node.go
-'  22  19  ~/.go/src/blockchain-simulator/src/blockchain/event.go
-'  37  5  ~/.go/src/blockchain-simulator/src/blockchain/event.go
-'  34  27  ~/.go/src/blockchain-simulator/src/blockchain/event.go
-'  10  34  ~/.go/src/blockchain-simulator/src/core/context.go
-'  31  6  ~/.go/src/blockchain-simulator/src/blockchain/event.go
-'  36  14  ~/.go/src/blockchain-simulator/src/blockchain/event.go
-'  1  0  ~/.go/src/blockchain-simulator/src/core/timeframe.go
-'  9  0  ~/.go/src/blockchain-simulator/src/core/context.go
-'  1  0  ~/.go/src/blockchain-simulator/src/blockchain/event.go
-'  9  16  ~/.go/src/blockchain-simulator/src/core/iterator.go
-'  32  0  ~/.go/src/blockchain-simulator/src/consensus/consensus.go
-'  9  0  ~/.go/src/blockchain-simulator/src/blockchain/blockchain.go
-'  1  11  ~/.go/src/blockchain-simulator/src/util/debug.go
-'  12  0  ~/.go/src/blockchain-simulator/src/core/iterator.go
-'  5  0  ~/.go/src/blockchain-simulator/src/blockchain/block.go
-'  1  0  ~/.go/src/blockchain-simulator/src/blockchain/block.go
-'  225  0  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  239  2  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  50  16  ~/.go/src/blockchain-simulator/src/core/blockchain_test.go
-'  8  11  ~/.go/src/blockchain-simulator/src/core/blockchain_test.go
-'  16  0  ~/.go/src/blockchain-simulator/src/core/node.go
-'  9  0  ~/.go/src/blockchain-simulator/src/context.go
-'  7  0  ~/.go/src/blockchain-simulator/src/timeframe.go
-'  15  0  ~/.go/src/blockchain-simulator/src/context.go
-'  47  1  ~/.go/src/blockchain-simulator/src/iterator.go
-'  5  0  ~/.go/src/blockchain-simulator/src/iterator.go
-'  19  0  ~/.go/src/blockchain-simulator/src/core/node.go
-'  1  0  ~/.go/src/blockchain-simulator/src/core/transaction.go
-'  38  1  ~/.go/src/blockchain-simulator/src/core/node.go
-'  46  1  ~/.go/src/blockchain-simulator/src/core/event.go
-'  25  1  ~/.go/src/blockchain-simulator/src/core/blockchain_test.go
-'  21  2  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  12  1  ~/.go/src/blockchain-simulator/src/core/block.go
-'  3  6  ~/.go/src/blockchain-simulator/src/iterator.go
-'  284  0  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  295  2  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  74  3  ~/.go/src/blockchain-simulator/src/core/node.go
-'  23  12  ~/.go/src/blockchain-simulator/src/network/network.go
-'  28  15  ~/.go/src/blockchain-simulator/src/network/network_test.go
-'  34  6  ~/.go/src/blockchain-simulator/src/core/node.go
-'  1  0  ~/.go/src/blockchain-simulator/src/core/simulator.go
-'  48  0  ~/.go/src/blockchain-simulator/src/core/node.go
-'  44  0  ~/.go/src/blockchain-simulator/src/iterator.go
-'  42  21  ~/.go/src/blockchain-simulator/src/iterator.go
-'  382  0  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  391  2  ~/.go/src/blockchain-simulator/NERD_tree_1
-'  28  0  ~/.go/src/blockchain-simulator/src/core/event.go
-'  32  17  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  16  23  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  15  3  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  19  15  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  24  18  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  43  1  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  20  9  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  16  1  ~/.go/src/blockchain-simulator/src/iterator.go
-'  1  11  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  7  1  ~/.go/src/blockchain-simulator/src/core/blockchain.go
-'  45  1  ~/.go/src/blockchain-simulator/src/core/block.go

# ë²„í¼ ëª©ë¡:
%~/.go/src/blockchain-simulator/src/blockchain/block.go	24	0
%~/.go/src/blockchain-simulator/src/config.json	1	0
%~/.go/src/ScamList/core/pow.go	57	34
%~/.go/src/blockchain-simulator/gopath_set.sh	1	33
%~/.go/src/blockchain-simulator/src/network/topology/subnet.go	1	0
%~/.go/src/blockchain-simulator/README.md	14	1
%~/.go/src/ScamList/core/block.go	69	0
%~/.go/src/github.com/ironjr/ScamList/cli/cli.go	40	13
%~/.go/src/github.com/ironjr/blockchain-simulator/src/network/topology/coord.go	9	6
%~/.go/src/github.com/ironjr/blockchain-simulator/src/network/topology/netnode.go	4	0
%~/.go/src/github.com/ironjr/blockchain-simulator/src/network/topology/subnet-gate.go	1	0
%~/.go/src/github.com/ironjr/ScamList/cli/cli_createblockchain.go	1	0
%~/.go/src/github.com/ironjr/ScamList/cli/cli_createwallet.go	1	0
%~/.go/src/github.com/ironjr/ScamList/core/block.go	1	11
%~/.go/src/github.com/ironjr/ScamList/core/chain.go	16	0
%~/.go/src/ScamList/cli/cli_printchain.go	12	0
%~/.go/src/ScamList/core/iterator.go	27	2
%~/.go/src/ScamList/README.md	34	0
%~/.go/src/blockchain-simulator/src/network/node.go	17	9
%~/.go/src/blockchain-simulator/src/network/test.go	35	2
%~/.go/src/blockchain-simulator/src/network/topology/netnode.go	1	0
%~/.go/src/blockchain-simulator/src/core/block.go	12	1
%~/.go/src/blockchain-simulator/src/core/blockchain.go	21	2
%~/.go/src/blockchain-simulator/src/core/merkle.go	120	0
%~/.go/src/blockchain-simulator/src/core/test.go	4	0
%~/.go/src/blockchain-simulator/src/core/transaction.go	1	0
%~/.go/src/ScamList/cli/cli.go	8	0
%~/.go/src/ScamList/cli/cli_createwallet.go	4	0
%~/.go/src/ScamList/main.go	1	0
%~/.go/src/ScamList/wallet_3000.dat	1	0
%~/.go/src/blockchain-simulator/src/network/topology/subnet-gate.go	4	8
%~/.go/src/blockchain-simulator/src/network/topology/coord.go	1	1
%~/.go/src/ScamList/cli/cli_send.go	25	1
%~/.go/src/ScamList/core/tx.go	3	0
%~/.go/src/ScamList/core/wallet.go	47	4
%~/.go/src/blocksim-old/README.md	1	0
%~/.go/src/blockchain-simulator/src/network/topology/subnet_gate.go	14	0
%~/.go/src/blockchain-simulator/src/network/config.json	1	0
%~/.go/src/blockchain-simulator/topology.go	1	0
%~/.go/src/blockchain-simulator/src/network/topology/topology.go	2	1
%~/.go/src/blockchain-simulator/src/network/topology/net_node.go	14	0
%~/.go/src/blockchain-simulator/src/network/topology/netnode_helper.go	1	8
%~/.go/src/blockchain-simulator/src/network/topology/subnetgate_helper.go	1	0
%~/.go/src/blockchain-simulator/src/network/topology/subnetgate.go	14	2
%~/.go/src/blockchain-simulator/src/network/topology/coord_model.go	1	0
%~/.go/src/blockchain-simulator/pkg/linux_amd64/github.com/OneOfOne/xxhash.a	55	0
%~/.go/src/blockchain-simulator/src/network/topology/coord_policy.go	52	0
%~/.go/src/blockchain-simulator/src/network/topology/subnetgate_policy.go	1	11
%~/.go/src/blockchain-simulator/src/network/topology/subnet_policy.go	1	0
%~/.go/src/blockchain-simulator/src/network/topology/delay_model.go	1	0
%~/.go/src/blockchain-simulator/src/network/topology/thruput_model.go	1	0
%~/.go/src/blockchain-simulator/src/network/topology/delaystat_model.go	3	1
%~/.go/src/blockchain-simulator/src/network/topology/packet_model.go	1	2
%~/.go/src/ScamList/core/chain.go	47	0
%~/.go/src/blockchain-simulator/src/util/debug.go	1	11
%~/.go/src/blockchain-simulator/src/network/topology/subnet_model.go	1	0
%~/.go/src/blockchain-simulator/src/network/topology/rtt.go	1	0
%~/.go/src/blockchain-simulator/src/network/event.go	6	0
%~/.go/src/blockchain-simulator/src/network/network.go	94	2
%~/.go/src/blockchain-simulator/src/network/network_test.go	28	15
%~/.go/src/blockchain-simulator/src/network/topology/topology_test.go	84	2
%~/.go/src/blockchain-simulator/src/core/simulator/node.go	1	0
%~/.go/src/blockchain-simulator/src/core/simulator/simulator.go	4	1
%~/.go/src/blockchain-simulator/src/core/blockchain/block.go	15	0
%~/.go/src/blockchain-simulator/src/core/blockchain/blockchain.go	2	0
%~/.go/src/blockchain-simulator/src/core/blockchain/transaction.go	6	0
%~/.go/src/blockchain-simulator/src/debug/test/topology_test_conf.json	26	4
%~/.go/src/blockchain-simulator/src/core/node.go	16	0
%~/.go/src/blockchain-simulator/src/core/simulator.go	8	0
%~/.go/src/blockchain-simulator/src/network/topology/debug/test/topology_test_conf.json	23	8
%~/.go/src/blockchain-simulator/src/core/simulator_test.go	1	0
%~/.go/src/blockchain-simulator/src/consensus/consensus.go	6	33
%~/.go/src/blockchain-simulator/src/iterator.go	47	1
%~/.go/src/blockchain-simulator/src/topology/topology.go	94	30
%~/.go/src/blockchain-simulator/src/topology/coord.go	7	0
%~/.go/src/blockchain-simulator/src/core/event.go	46	1
%~/.go/src/blockchain-simulator/src/topology/topology_test.go	43	18
%~/.go/src/blockchain-simulator/src/topology/coord_model.go	19	0
%~/.go/src/blockchain-simulator/src/core/blockchain_test.go	50	16
%~/.go/src/blockchain-simulator/src/context.go	1	0
%~/.go/src/blockchain-simulator/src/timeframe.go	12	31
%~/.go/src/blockchain-simulator/src/core/iterator.go	1	11
%~/.go/src/blockchain-simulator/src/blockchain/blockchain.go	9	0
%~/.go/src/blockchain-simulator/src/blockchain/node.go	70	0
%~/.go/src/blockchain-simulator/src/blockchain/event.go	7	0
%~/.go/src/blockchain-simulator/src/core/context.go	1	11
%~/.go/src/blockchain-simulator/src/core/timeframe.go	14	0
%~/.go/src/blockchain-simulator/src/network/coord.go	6	0
%~/.go/src/blockchain-simulator/src/network/delay_model.go	53	0
%~/.go/src/blockchain-simulator/src/txgen/model.go	20	18
%~/.go/src/blockchain-simulator/src/txgen/txgen.go	1	12

# íŒŒì¼ë‚´ì˜ ë§ˆí¬ íˆìŠ¤í† ë¦¬ (ìƒˆê²ƒë¶€í„° ì˜¤ë˜ëœ ìˆœ):

> ~/.go/src/blockchain-simulator/src/txgen/model.go
	"	20	18
	^	48	1
	.	46	0
	+	10	8
	+	52	0
	+	54	0
	+	7	0
	+	1	0
	+	2	7
	+	12	29
	+	13	27
	+	17	24
	+	20	11
	+	19	11
	+	20	18
	+	23	5
	+	18	3
	+	23	37
	+	24	0
	+	31	19
	+	27	32
	+	28	11
	+	24	5
	+	28	11
	+	24	11
	+	28	13
	+	31	29
	+	32	0
	+	13	0
	+	14	8
	+	50	0
	+	32	0
	+	33	0
	+	38	39
	+	39	14
	+	34	35
	+	35	11
	+	36	1
	+	40	0
	+	42	62
	+	43	0
	+	45	43
	+	46	0
	+	47	0
	+	48	0
	+	43	0
	+	46	0
	z	26	0

> ~/.go/src/blockchain-simulator/src/txgen/txgen.go
	"	1	12
	^	1	13
	.	1	12
	+	1	12
	z	1	12

> ~/.go/src/blockchain-simulator/NERD_tree_1
	"	26	0
	^	611	18
	.	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0

> ~/.go/src/blockchain-simulator/src/network/network.go
	"	94	2
	.	11	0
	+	16	0
	+	15	0
	+	11	0
	z	57	0

> ~/.go/src/blockchain-simulator/src/network/delay_model.go
	"	53	0

> ~/.go/src/blockchain-simulator/src/core/simulator.go
	"	8	0

> ~/.go/src/blockchain-simulator/src/blockchain/event.go
	"	7	0
	^	40	1
	.	4	15
	+	36	14
	+	35	2
	+	34	27
	+	37	5
	+	4	15
	z	1	0

> ~/.go/src/blockchain-simulator/src/network/coord.go
	"	6	0

> ~/.go/src/blockchain-simulator/src/consensus/consensus.go
	"	6	33
	^	6	33
	.	6	0
	+	20	0
	+	21	0
	+	32	0
	+	3	8
	+	4	0
	+	8	0
	+	4	0
	+	7	17
	+	10	15
	+	6	0
	z	10	15

> ~/.go/src/blockchain-simulator/src/core/iterator.go
	"	1	11
	^	1	12
	.	1	11
	+	11	0
	+	1	11
	z	1	11

> ~/.go/src/blockchain-simulator/src/core/context.go
	"	1	11
	^	1	12
	.	1	11
	+	10	34
	+	1	11
	z	1	11

> ~/.go/src/blockchain-simulator/src/core/timeframe.go
	"	14	0

> ~/.go/NERD_tree_1
	"	4	0
	.	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0
	+	1	0

> ~/.go/src/blockchain-simulator/src/blockchain/block.go
	"	24	0

> ~/.go/src/blockchain-simulator/src/util/debug.go
	"	1	11
	^	1	12
	.	1	11
	+	1	11
	z	1	11

> ~/.go/src/blockchain-simulator/src/blockchain/blockchain.go
	"	9	0

> ~/.go/src/blockchain-simulator/src/blockchain/node.go
	"	70	0

> ~/.go/src/blockchain-simulator/src/core/blockchain_test.go
	"	50	16
	.	53	19
	+	4	0
	+	1	11
	+	10	15
	+	11	10
	+	12	8
	+	13	10
	+	20	19
	+	23	8
	+	24	10
	+	31	19
	+	34	8
	+	35	10
	+	42	19
	+	45	8
	+	46	10
	+	53	19
	z	50	16

> ~/.go/src/blockchain-simulator/src/core/block.go
	"	12	1
	^	40	27
	.	40	26
	+	1	11
	+	7	16
	+	9	20
	+	26	26
	+	1	0
	+	38	26
	+	26	0
	+	41	0
	+	22	8
	+	27	0
	+	26	3
	+	36	25
	+	37	10
	+	41	2
	+	51	2
	+	40	29
	+	47	19
	+	52	3
	+	37	0
	+	53	10
	+	51	51
	+	52	61
	+	40	26
	z	17	0

> ~/.go/src/blockchain-simulator/src/core/blockchain.go
	"	21	2
	^	21	27
	.	21	26
	+	38	0
	+	44	0
	+	35	0
	+	20	0
	+	13	0
	+	11	0
	+	15	20
	+	5	10
	+	10	27
	+	11	2
	+	12	10
	+	19	28
	+	20	0
	+	19	0
	+	24	18
	+	19	15
	+	15	37
	+	17	1
	+	17	1
	+	17	0
	+	16	24
	+	25	14
	+	31	0
	+	37	19
	+	21	26
	z	21	26
