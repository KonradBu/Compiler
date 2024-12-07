// Generated from /home/konradburgi/Documents/hhu/gitCode/compilerbau/antlr/kbgrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class kbgrammarLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		NAMESPACE=1, USING_BLOCK=2, CLASS=3, MAIN=4, FUNC_BLOCK=5, NAME=6;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"NAMESPACE", "USING_BLOCK", "CLASS", "MAIN", "FUNC_BLOCK", "NAME"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "NAMESPACE", "USING_BLOCK", "CLASS", "MAIN", "FUNC_BLOCK", "NAME"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public kbgrammarLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "kbgrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0006U\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004"+
		"\u0001\u0004\u0001\u0005\u0001\u0005\u0000\u0000\u0006\u0001\u0001\u0003"+
		"\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006\u0001\u0000\u0000T"+
		"\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000"+
		"\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000"+
		"\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0001"+
		"\r\u0001\u0000\u0000\u0000\u0003\u001c\u0001\u0000\u0000\u0000\u0005&"+
		"\u0001\u0000\u0000\u0000\u00079\u0001\u0000\u0000\u0000\tQ\u0001\u0000"+
		"\u0000\u0000\u000bS\u0001\u0000\u0000\u0000\r\u000e\u0005n\u0000\u0000"+
		"\u000e\u000f\u0005a\u0000\u0000\u000f\u0010\u0005m\u0000\u0000\u0010\u0011"+
		"\u0005e\u0000\u0000\u0011\u0012\u0005s\u0000\u0000\u0012\u0013\u0005p"+
		"\u0000\u0000\u0013\u0014\u0005a\u0000\u0000\u0014\u0015\u0005c\u0000\u0000"+
		"\u0015\u0016\u0005e\u0000\u0000\u0016\u0017\u0001\u0000\u0000\u0000\u0017"+
		"\u0018\u0003\u000b\u0005\u0000\u0018\u0019\u0005{\u0000\u0000\u0019\u001a"+
		"\u0003\u0005\u0002\u0000\u001a\u001b\u0005}\u0000\u0000\u001b\u0002\u0001"+
		"\u0000\u0000\u0000\u001c\u001d\u0005u\u0000\u0000\u001d\u001e\u0005s\u0000"+
		"\u0000\u001e\u001f\u0005i\u0000\u0000\u001f \u0005n\u0000\u0000 !\u0005"+
		"g\u0000\u0000!\"\u0001\u0000\u0000\u0000\"#\u0003\u000b\u0005\u0000#$"+
		"\u0005;\u0000\u0000$%\u0003\u0003\u0001\u0000%\u0004\u0001\u0000\u0000"+
		"\u0000&\'\u0005p\u0000\u0000\'(\u0005u\u0000\u0000()\u0005b\u0000\u0000"+
		")*\u0005l\u0000\u0000*+\u0005i\u0000\u0000+,\u0005c\u0000\u0000,-\u0001"+
		"\u0000\u0000\u0000-.\u0005c\u0000\u0000./\u0005l\u0000\u0000/0\u0005a"+
		"\u0000\u000001\u0005s\u0000\u000012\u0005s\u0000\u000023\u0001\u0000\u0000"+
		"\u000034\u0003\u000b\u0005\u000045\u0005{\u0000\u000056\u0003\u0007\u0003"+
		"\u000067\u0003\t\u0004\u000078\u0005}\u0000\u00008\u0006\u0001\u0000\u0000"+
		"\u00009:\u0005p\u0000\u0000:;\u0005u\u0000\u0000;<\u0005b\u0000\u0000"+
		"<=\u0005l\u0000\u0000=>\u0005i\u0000\u0000>?\u0005c\u0000\u0000?@\u0001"+
		"\u0000\u0000\u0000@A\u0005s\u0000\u0000AB\u0005t\u0000\u0000BC\u0005a"+
		"\u0000\u0000CD\u0005t\u0000\u0000DE\u0005i\u0000\u0000EF\u0005c\u0000"+
		"\u0000FG\u0001\u0000\u0000\u0000GH\u0005v\u0000\u0000HI\u0005o\u0000\u0000"+
		"IJ\u0005i\u0000\u0000JK\u0005d\u0000\u0000KL\u0001\u0000\u0000\u0000L"+
		"M\u0005m\u0000\u0000MN\u0005a\u0000\u0000NO\u0005i\u0000\u0000OP\u0005"+
		"n\u0000\u0000P\b\u0001\u0000\u0000\u0000QR\u0001\u0000\u0000\u0000R\n"+
		"\u0001\u0000\u0000\u0000ST\u0001\u0000\u0000\u0000T\f\u0001\u0000\u0000"+
		"\u0000\u0001\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}