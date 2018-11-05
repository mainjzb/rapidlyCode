#include <QtCore/QCoreApplication>
#include <iostream>
#include <QFile>
#include <QTextStream>
#include <QTextCodec>
#include <QDebug>
#include <QDateTime>
#include <QDate>
#include "cmdline.h"


struct sRule
{
	QRegExp before;
	QString after;
};


void XParseRule( QString& ruleFile_s, QVector<sRule*>& rules );
QString * XReplace( QString * content, const QVector<sRule*>& rules );
void SaveToFile( QTextStream & outStream, QString & content, const bool& AD, const QString & adFile );
bool XPrint( const char* );
void CoverDateTimeToBeijin( QString & utcTime );
void OptimizeContent( QString & content );



int main( int argc, char *argv[] )
{
	QCoreApplication *a = new QCoreApplication( argc, argv );

	cmdline::parser cmdopt;
	cmdopt.add<std::string>( "src", 's', "in file", false, "./a.html" );
	cmdopt.add<std::string>( "dest", 'd', "out file", false, "./b.html" );
	cmdopt.add<std::string>( "rule", 'r', "rule file", false, "./rule.md" );
	cmdopt.add( "noad", '\0', "don't add ad to out File" );
	cmdopt.add<std::string>( "adpath", '\0', "ad file", false, "./Tools/ad.html" );

	cmdopt.add( "help", 0, "print this message" );
	cmdopt.footer( "filename ..." );
	cmdopt.set_program_name( "IR" );
	bool ok1 = cmdopt.parse( argc, argv );

	if ( argc == 1 || cmdopt.exist( "help" ) )
	{
		std::cout << cmdopt.usage();
	}

	QString	ruleFile_s = QString::fromStdString( cmdopt.get<std::string>( "rule" ) );
	QString	srcFile_s = QString::fromStdString( cmdopt.get<std::string>( "src" ) );
	QString	destFile_s = QString::fromStdString( cmdopt.get<std::string>( "dest" ) );
	bool	adFlag = !cmdopt.exist( "noad" );
	QString	adPath = QString::fromStdString( cmdopt.get<std::string>( "adpath" ) );


	//set UTF-8
	QTextCodec::setCodecForLocale( QTextCodec::codecForName( "UTF8" ) );
	size_t ruleCount = cmdopt.rest().size();
	int iRule = 0;

	QFile srcFile( srcFile_s );
	QFile destFile( destFile_s );
	if ( !srcFile.open( QFile::ReadWrite ) || !destFile.open( QFile::ReadWrite ) )
	{
		return false;
	}

	QTextStream inStream( &srcFile );
	QTextStream outStream( &destFile );
	inStream.setCodec( "UTF-8" );
	outStream.setCodec( "UTF-8" );

	QString * content = new QString( u8"" );
	*content = inStream.readAll();
	OptimizeContent( *content );

	QVector<sRule*> rules;
	XParseRule(ruleFile_s, rules);

	content = XReplace( content, rules );

	SaveToFile( outStream, *content, adFlag, adPath );

	XPrint( "Success!" );

#ifdef QT_DEBUG
	a->exec();
#else
	a->exit();
#endif

	return 0;
}

QString * XReplace( QString * content, const QVector<sRule*>& rules )
{
	if ( content == nullptr )
	{
		XPrint( "[XReplace]: content is Null" );
		return nullptr;
	}

	QString * result = new QString();

#ifdef QT_DEBUG
	QTextStream out( stdout );
	//qDebug() << src;
#endif


	int start = 0, end = 1;
	const int contentLength = content->length();
	while ( start < contentLength )
	{
		if ( content->at( start ) == '<' )
		{

			end = content->indexOf( '>', start + 1 );
			while ( end + 1 < contentLength && content->at( end + 1 ) == '<' )
			{
				end = content->indexOf( '>', end + 2 );
			}
			result->append( content->mid( start, end + 1 - start ) );
			//content.remove( 0, end + 1 );
			start = end + 1;
		}
		else
		{
			end = content->indexOf( '<', start + 1 );
			if ( end == -1 )
			{
				end = contentLength;
			}
			QString tmpString = content->mid( start, end - start );
			tmpString = tmpString.trimmed();
			start = end;

			if ( tmpString.leftRef( 3 ) == "UTC" || tmpString.leftRef( 4 ) == "AEDT")
			{
				CoverDateTimeToBeijin( tmpString );
			}

			if ( tmpString.length() == 0 ||
				tmpString.left( 3 ) == "PDT" || 
				tmpString.left( 3 ) == "PST" || 
				tmpString.left( 3 ) == "EDT" || 
				tmpString.left( 3 ) == "EST" || 
				tmpString.left( 4 ) == "CEST" ||
				tmpString.left( 3 ) == "CET" )
			{
				continue;
			}

			for ( const sRule* rule : rules )
			{
				tmpString.replace( rule->before, rule->after );
			}

			result->append( tmpString );
		}
	}


	delete content;
	return result;
}

bool XPrint( const char* s )
{
	//QTextStream out( stdout );
	//out << s << endl;
	qDebug() << s;
	return true;
}

void XParseRule( QString& ruleFile_s, QVector<sRule*>& rules )
{
	QFile ruleFile( ruleFile_s );
	if ( !ruleFile.open( QFile::ReadWrite ) )
	{
		XPrint( "[XRule2Content]: ruleFile open error" );
		Q_ASSERT(false);
	}
	QString ruleFilePath = ruleFile_s.left( ruleFile_s.lastIndexOf( "/" ) + 1 );


	QTextStream replaceStream( &ruleFile );
	replaceStream.setCodec( "UTF-8" );

	while ( !replaceStream.atEnd() )
	{
		QString rule = replaceStream.readLine().trimmed();
		if ( rule.isEmpty() || ( rule.length() > 0 && ( rule.at( 0 ) == "#" || rule.at( 0 ) == "-" ) ) )
			continue;

		if ( rule.length() > 0 && rule.left( 4 ) == R"(<!--)" )
			continue;

		if ( rule.length() > 0 && rule.at( 0 ) == '@' )
		{
			rule.remove( 0, 1 ).trimmed();
			if ( rule.left( 6 ) == "import" )
			{
				int start = rule.indexOf( '"' );
				int end = rule.lastIndexOf( '"' );
				QString importRulePath = ruleFilePath + rule.mid( start + 1, end - start - 1 ).trimmed();
				XParseRule( importRulePath, rules );
				continue;
			}
		}


		QStringList ruleList = rule.split( "==" );

		Q_ASSERT( ruleList.length() == 2 );

		sRule * newRule = new sRule();
		newRule->before = QRegExp( ruleList.at( 0 ) );
		newRule->after = ruleList.at( 1 );
		rules.push_back( newRule );
		
		if ( ruleList.at( 0 ).at( 0 ).isLower() )
		{
			QString rule_src = ruleList.at( 0 );
			QString rule_dst = ruleList.at( 1 );

			QString::iterator it = rule_src.begin();
			*it = ( *it ).toUpper();

			sRule * newRule2 = new sRule();
			newRule2->before = QRegExp( rule_src );
			newRule2->after = rule_dst;
			rules.push_back( newRule2 );
		}
	}

	ruleFile.close();
}

void SaveToFile( QTextStream & outStream, QString & content, const bool& AD, const QString & adFilePath)
{
	if ( AD == true )
	{
		QString adString;
		QFile adFile( adFilePath );
		if ( adFile.open( QFile::ReadWrite ) )
		{
			adString = adFile.readAll();
		}
		int ans = content.lastIndexOf( R"(</div>)" );
		if ( ans != -1 )
		{
			content.remove( ans, 6 );
			outStream << content;
			outStream.flush();
			outStream << adString;
			outStream << R"(</div>)";
			outStream.flush();
		}
		else
		{
			outStream << content;
			outStream << adString;
			outStream.flush();
		}
	}
	else
	{
		outStream << content;
		outStream.flush();
	}
}

void CoverDateTimeToBeijin(QString & utcTime)
{
	QString result1, result2;

	if ( utcTime.leftRef( 3 ) == "UTC" )
	{
		utcTime.remove( 0, 4 ).trimmed();


		QStringList tmpstring = utcTime.split( QString::fromWCharArray( L"–" ) );
		if ( tmpstring.length() < 2 )
		{
			tmpstring = utcTime.split( '-' );
		}
		if ( tmpstring.length() != 2 )
		{
			return;
		}

		QString startTime = tmpstring.at( 0 ).trimmed();
		QString endTime = tmpstring.at( 1 ).trimmed();

		QLocale loc = QLocale( QLocale::English );
		QDateTime date1 = loc.toDateTime( startTime, "MMMM d 'at' h:mm AP" );
		QDateTime date2 = loc.toDateTime( endTime, "MMMM d 'at' h:mm AP" );
		QDate currentDate = QDate::currentDate();


		if ( !date1.isValid() && !date2.isValid() )
		{
			return;
		}
		if ( date1.isValid() )
		{
			date1 = date1.addSecs( qint64( 8 * 3600 ) );
			date1.setDate( QDate( currentDate.year(), date1.date().month(), date1.date().day() ) );
			result1 = date1.toString( L"M月d日 ap h:mm" );
		}
		else
		{
			result1 = startTime;
		}

		if ( date2.isValid() )
		{
			date2 = date2.addSecs( qint64( 8 * 3600 ) );
			date2.setDate( QDate( currentDate.year(), date2.date().month(), date2.date().day() ) );
			result2 = date2.toString( L"M月d日 ap h:mm" );
		}
		else
		{
			result2 = endTime;
		}
	}

	if ( utcTime.leftRef( 4 ) == "AEDT" )
	{
		int index = utcTime.indexOf( ':' );
		if ( index < 0 )
		{
			Q_ASSERT( "AEDT time can't find ':' " );
		}
		int index2 = utcTime.indexOf( '+' );
		int rightTime = 0;
		if ( index2 < 0 )
		{
			Q_ASSERT( "AEDT time can't find '+' " );
		}
		if ( utcTime.midRef( index2 + 1, 2 ) == "10" )
		{
			rightTime = -2;
		}
		if ( utcTime.midRef( index2 + 1, 2 ) == "11" )
		{
			rightTime = -3;
		}

		utcTime.remove( 0, index + 1 );
		QStringList tmpstring = utcTime.split( QString::fromWCharArray( L"–" ) );
		if ( tmpstring.length() < 2 )
		{
			tmpstring = utcTime.split( '-' );
		}
		if ( tmpstring.length() != 2 )
		{
			return;
		}
		QString startTime = tmpstring.at( 0 ).trimmed();
		QString endTime = tmpstring.at( 1 ).trimmed();
		startTime = startTime.replace( "&nbsp;", " " );
		endTime = endTime.replace( "&nbsp;", " " );
		startTime = startTime.replace( "  ", " " );
		endTime = endTime.replace( "  ", " " );

		QLocale loc = QLocale( QLocale::English );
		QDateTime date1 = loc.toDateTime( startTime, "dddd, MMMM d, yyyy h:mm AP" );
		QDateTime date2 = loc.toDateTime( endTime, "dddd, MMMM d, yyyy h:mm AP" );


		if ( !date1.isValid() && !date2.isValid() )
		{
			return;
		}
		if ( date1.isValid() )
		{
			date1 = date1.addSecs( qint64( rightTime * 3600 ) );
			result1 = date1.toString( L"yyyy年 M月d日 ap h:mm" );
		}
		else
		{
			result1 = startTime;
		}

		if ( date2.isValid() )
		{
			date2 = date2.addSecs( qint64( rightTime * 3600 ) );
			result2 = date2.toString( L"yyyy年 M月d日 ap h:mm" );
		}
		else
		{
			result2 = endTime;
		}

	}
	utcTime.clear();
	utcTime = QString(u8"北京时间：") + result1 + "  -  " + result2;
}


void OptimizeContent(QString & content)
{
	if ( content == nullptr )
	{
		XPrint( "[OptimizeContent]: content is Null" );
		return ;
	}

	content = content.replace( R"(</strong><strong>)", "" );
	content = content.replace( QRegExp(R"(</strong>\s<strong>)"), " " );
	content = content.replace( R"(</strong>&nbsp;<strong>)", " " );



	int start = 0, end = 0;
	while (true)
	{
		start = content.indexOf( R"(<strong>)", start );
		if ( start == -1 )
			break;

		int index1 = content.indexOf( R"(</strong>)", start );
		if ( index1 == -1 )
			XPrint( "[OptimizeContent] can't find </strong>" );
		int index2 = content.indexOf( R"(<strong>)", start + 7 );
		if ( index2 == -1 )
			break;

		if ( index1 > index2 )
		{

			if ( content.mid( index1 - 1, 1 ) == " " && content.mid( index1 + 9, 1 ) == " " )
			{
				content.remove( index1, 10 );
			}
			else
			{
				content.remove( index1, 9 );
			}

			if ( content.mid( index2 - 1, 1 ) == " " && content.mid( index2 + 8, 1 ) == " " )
			{
				content.remove( index2, 9 );
			}
			else
			{
				content.remove( index2, 8 );
			}
			start = index1;
		}
		else
		{
			start = index2;
		}

	}
}