#include <QtCore/QCoreApplication>
#include <iostream>
#include <QFile>
#include <QTextStream>
#include <QTextCodec>
#include <QDebug>
#include "cmdline.h"

QString * XRule2Content( QString ruleFile_s, QString * content );
QString * XReplace( QString * content, const QStringList& ruleList );
bool XPrint( const char* s );



int main( int argc, char *argv[] )
{
	QCoreApplication *a = new QCoreApplication( argc, argv );

	cmdline::parser cmdopt;
	cmdopt.add<std::string>( "src", 's', "src file", false, "./a.html" );
	cmdopt.add<std::string>( "dest", 'd', "dest file", false, "./b.html" );
	cmdopt.add<std::string>( "rule", 'r', "rule file", false, "./rule.md" );
	cmdopt.add( "help", 0, "print this message" );
	cmdopt.footer( "filename ..." );
	cmdopt.set_program_name( "IR" );
	bool ok1 = cmdopt.parse( argc, argv );

	if ( argc == 1 || cmdopt.exist( "help" ) )
	{
		std::cout << cmdopt.usage();
	}

	QString ruleFile_s = QString::fromStdString( cmdopt.get<std::string>( "rule" ) );
	QString srcFile_s = QString::fromStdString( cmdopt.get<std::string>( "src" ) );
	QString destFile_s = QString::fromStdString( cmdopt.get<std::string>( "dest" ) );

	//set UTF-8
	QTextCodec::setCodecForLocale( QTextCodec::codecForName( "UTF8" ) );
	int ruleCount = cmdopt.rest().size();
	int iRule = 0;

	QFile srcFile( srcFile_s );
	QFile destFile( destFile_s );
	if ( !srcFile.open( QFile::ReadWrite ) || !destFile.open( QFile::ReadWrite ) )
	{
		return false;
	}

	QTextStream stream( &srcFile );
	QTextStream outStream( &destFile );
	stream.setCodec( "UTF-8" );
	outStream.setCodec( "UTF-8" );

	QString * content = new QString( u8"" );
	*content = stream.readAll();

	while ( true )
	{
		content = XRule2Content(ruleFile_s, content);
		if ( iRule != ruleCount )
		{
			ruleFile_s = QString::fromStdString( cmdopt.rest()[iRule++] );
		}
		else
		{
			break;
		}
	}

	outStream << *content;
	outStream.flush();

	XPrint( "Success!" );

#ifdef QT_DEBUG
	a->exec();
#else
	a->exit();
#endif

	return 0;
}

QString * XReplace( QString * content, const QStringList & ruleList )
{
	if ( content == nullptr )
	{
		XPrint( "[XReplace]: content is Null" );
		return nullptr;
	}


	QString * result = new QString();
	QRegExp src = QRegExp( ruleList.at( 0 ) );
	QString dest = ruleList.at( 1 );


#ifdef QT_DEBUG
	QTextStream out( stdout );
	//out << src << "==" << dest << endl;
	qDebug() << src;
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
			//content.remove( 0, end );
			start = end;

			tmpString.replace( src, dest );
			result->append( tmpString );
		}
	}

	delete content;
	return result;
}

bool XPrint( const char* s )
{
	QTextStream out( stdout );
	out << s << endl;
	qDebug() << s;
	return true;
}

QString * XRule2Content( QString ruleFile_s, QString * content )
{
	QFile ruleFile( ruleFile_s );
	if ( !ruleFile.open( QFile::ReadWrite ) )
	{
		XPrint( "[XRule2Content]: ruleFile open error" );
		Q_ASSERT(false);
		return nullptr;
	}
	QString ruleFilePath = ruleFile_s.left( ruleFile_s.lastIndexOf( "/" ) + 1 );


	QTextStream replaceStream( &ruleFile );
	replaceStream.setCodec( "UTF-8" );


	while ( !replaceStream.atEnd() )
	{
		QString rule = replaceStream.readLine().trimmed();
		if ( rule.isEmpty() || ( rule.length() > 0 && ( rule.at( 0 ) == "#" || rule.at( 0 ) == "-" ) ) )
			continue;

		if ( rule.length() > 0 && rule.at( 0 ) == '@' )
		{
			rule.remove( 0, 1 ).trimmed();
			if ( rule.left( 6 ) == "import" )
			{
				int start = rule.indexOf( '"' );
				int end = rule.lastIndexOf( '"' );
				QString importRulePath = ruleFilePath + rule.mid( start + 1, end - start - 1 ).trimmed();
				content = XRule2Content( importRulePath, content );
				continue;
			}
		}


		QStringList ruleList = rule.split( "==" );

		Q_ASSERT( ruleList.length() == 2 );

		content = XReplace( content, ruleList );
		if ( ruleList.at( 0 ).at( 0 ).isLower() )
		{
			QString rule_src = ruleList.at( 0 );
			QString rule_dst = ruleList.at( 1 );

			QString::iterator it = rule_src.begin();
			*it = ( *it ).toUpper();
			QStringList ruleList2;
			ruleList2.append( rule_src );
			ruleList2.append( rule_dst );
			content = XReplace( content, ruleList2 );
		}
	}


	ruleFile.close();
	return content;
}